package easyQueue

import (
	"fmt"
	"runtime/debug"
	"sync/atomic"
	"time"
)

const (
	maxWorkLen  = 200 //堆积go 数量
	minWorkLen  = 20  //最小go 数量
	growStepLen = 20  //每次增加GO 数量

	chanQueueLen = 10000 //？？
)

func init() {
	initWorkers(minWorkLen)
}

//work 处理队列
type ParamFunc struct {
	Param    interface{}
	Func     func(inParam interface{}) (err error)
	CallBack func(inParam interface{}, func_err error) error //default nil
}

type workPool struct {
	closed bool
	size   int32 //all go
	works  chan ParamFunc
	cancel chan struct{}
}

var wp *workPool

func initWorkers(workerNum int) {

	if workerNum < minWorkLen || workerNum > maxWorkLen {
		workerNum = minWorkLen
	}
	wp = &workPool{
		closed: false,
		size:   0,
		works:  make(chan ParamFunc, chanQueueLen),
		cancel: make(chan struct{}),
	}
	for i := 0; i < workerNum; i++ {
		go wp.startWorker()
	}
}

func (wp *workPool) startWorker() {
	atomic.AddInt32(&wp.size, 1)
	defer func() {
		atomic.AddInt32(&wp.size, -1)
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	t := time.NewTimer(30 * time.Second)
	defer t.Stop()

	now := time.Now().Unix()
	for {
		select {
		case entity, ok := <-wp.works:
			if !ok {
				return
			}
			now = time.Now().Unix()
			if entity.Func != nil {
				err := entity.Func(entity.Param)
				if entity.CallBack != nil {
					entity.CallBack(entity.Param, err)
				}
			}
		case <-t.C: //缩容
			t.Reset(30 * time.Second)
			if time.Now().Unix()-now > 5*60 { //5分钟空闲
				if atomic.LoadInt32(&wp.size) > minWorkLen {
					//xlog.LevelLogfn(xlog.WARN, "scale work pool idle....work size=%d", wp.size)
					fmt.Printf("scale work pool idle....work size=%d", wp.size)
					return
				}
			}
		case <-wp.cancel:
			return
		}
	}
}

func Stop() {
	if wp == nil {
		return
	}
	if wp.closed {
		return
	}
	size := atomic.LoadInt32(&wp.size)
	for i := 0; i < int(size); i++ {
		wp.cancel <- struct{}{}
	}
	wp.closed = true
	close(wp.works)
	close(wp.cancel)
}

func SendQueue(entity ParamFunc) bool {
	if wp == nil {
		//xlog.LevelLogfn(xlog.WARN, "SendQueue work pool not init.....")
		fmt.Printf("SendQueue work pool not init.....\n")
		return false
	}
	if wp.closed {
		//xlog.LevelLogfn(xlog.WARN, "SendQueue work pool closed.....")
		fmt.Printf("SendQueue work pool closed.....\n")
		return false
	}
	//堆积到maxWorkLen的倍数,开始扩容GO 数量
	if wp.size < minWorkLen || (wp.size <= maxWorkLen-growStepLen && int32(len(wp.works)-maxWorkLen)/maxWorkLen > (wp.size-minWorkLen)/growStepLen) {
		for i := 0; i < growStepLen; i++ {
			go wp.startWorker()
		}
		time.Sleep(50 * time.Millisecond)
		//xlog.LevelLogfn(xlog.INFO, "SendQueue work pool growStepLen....work size=%d,len=%d", wp.size, len(wp.works))
		fmt.Printf("SendQueue work pool growStepLen....work size=%d,len=%d\n", wp.size, len(wp.works))
	}
	select {
	case wp.works <- entity:
		return true
	default: //如果已经满了--直接退出，非阻塞
		//xlog.LevelLogfn(xlog.WARN, "SendQueue work pool busy....work size=%d,len=%d,entity:%+v", wp.size, len(wp.works), entity)
		fmt.Printf("SendQueue work pool busy....work size=%d,len=%d,entity:%+v\n", wp.size, len(wp.works), entity)
		//qywechat.SendQyWechat(qywechat.AlarmTypeInternalSrv, []string{fmt.Sprintf("SendQueue work pool busy....work size=%d,len=%d,entity:%+v", wp.size, len(wp.works), entity)}, global.QYWechatToBackOnly...)
		return false
	}
}
