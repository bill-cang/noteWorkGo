package main

import (
	"encoding/json"
	"fmt"
	m_struct "noteWork/基础类型/结构体"
	"reflect"
	"strings"
	"sync"
)

func main() {
	xiaoHei := &m_struct.Dog{
		Animal: m_struct.Animal{
			Name: "xiaoHei",
			Age:  3,
		},
		Trait: "love",
	}

	ain := &m_struct.Animal{
		Name: "",
		Age:  0,
	}

}

/*====================实参、形参：*/

type Foo struct {
	key    string
	option Option
	// ...
}
type Option struct {
	num int
	str string
}
type ModOption func(option *Option)

func New(key string, modOption ModOption) *Foo {
	option := Option{
		num: 100,
		str: "hello",
	}
	modOption(&option)
	return &Foo{
		key:    key,
		option: option,
	}
}

func WithNum(num int) ModOption {
	return func(option *Option) {
		option.num = num
	}
}
func WithStr(str string) ModOption {
	return func(option *Option) {
		option.str = str
	}
}

// ...
func main5() {
	foo := New("iamkey", func(option *Option) {
		// 调用方只设置 num
		option.num = 200
	})

	fmt.Printf("foo =%+v", foo)
}

/*====================实参、形参：*/
//个人活动热信息
type WinCashPersonalMsg struct {
	EndTime    int64    `protobuf:"varint,1,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	ObtainCash int32    `protobuf:"varint,2,opt,name=obtain_cash,json=obtainCash,proto3" json:"obtain_cash,omitempty"`
	DeviceIds  []string `protobuf:"bytes,4,rep,name=device_ids,json=deviceIds,proto3" json:"device_ids,omitempty"`
}

func main4() {
	x := [3]int{1, 2, 3}

	func(x [3]int) {
		// 此时x作为形参，只是对参数X复制额一份，是两个不同的参数。地址不在一样。
		x[0] = 7
	}(x)
	fmt.Printf("1->x =%+v\n", x) //prints [7 2 3]
	func(x *[3]int) {
		x[0] = 7
	}(&x)
	fmt.Printf("2->x =%+v\n", x) //prints [7 2 3]

}

func GetInvPersonalCache(uid string) (per *WinCashPersonalMsg, err error) {
	/*此段代码注释，当函数内部发生错误时，返回per是nil,空指针*/
	//per = &WinCashPersonalMsg{}
	return
}

/*====================反射获取结构体属性*/

type resume struct {
	Name   string `json:"name" doc:"我的名字"`
	age    int    `json:"age"`
	height int    `json:"HEIGHT"`
}

func findDoc(stru interface{}) map[string]string {
	t := reflect.TypeOf(stru).Elem()
	doc := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
	}

	return doc

}

func main3() {
	var stru resume
	doc := findDoc(&stru)
	for k, v := range doc {
		fmt.Printf("k = %s ,v = %v\n", k, v)
	}

}

/*====================redis  set 结合存入结构体数据，必先转换为字节在转换为string存入，取出时有字符串转换为字节反序列化实现*/
type ActInviteWinner struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InvCount int32  `protobuf:"varint,2,opt,name=inv_count,json=invCount,proto3" json:"inv_count,omitempty"`
	WinCash  int32  `protobuf:"varint,3,opt,name=win_cash,json=winCash,proto3" json:"win_cash,omitempty"`
}

func main2() {

	var winner = &ActInviteWinner{
		Name:     "ckx",
		InvCount: 70,
		WinCash:  1000,
	}
	marshal, _ := json.Marshal(winner)
	var data = string(marshal)
	var win = &ActInviteWinner{}
	err := json.Unmarshal([]byte(data), win)
	if err != nil {
		fmt.Printf("Unmarshal err =%+v\n", err)
		return
	}
	fmt.Printf("win =%+v\n", win)

}

/*================切片按下边存入数据*/
type TaskSpeed struct {
	Finished int32 `protobuf:"varint,1,opt,name=finished,proto3" json:"finished,omitempty"`
	Total    int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func main1() {
	tasks := make([]*TaskSpeed, 6)

	var wt = sync.WaitGroup{}
	wt.Add(5)
	go func() {
		tk := &TaskSpeed{
			Finished: 1,
			Total:    0,
		}
		tasks[0] = tk
		wt.Done()
	}()
	go func() {
		tk := &TaskSpeed{
			Finished: 2,
			Total:    0,
		}
		tasks[1] = tk
		wt.Done()
	}()
	go func() {
		tk := &TaskSpeed{
			Finished: 3,
			Total:    0,
		}
		tasks[2] = tk
		wt.Done()
	}()
	go func() {
		tk := &TaskSpeed{
			Finished: 4,
			Total:    0,
		}
		tasks[3] = tk
		wt.Done()
	}()
	/*	go func() {
		tk := &TaskSpeed{
			Finished: 5,
			Total:    0,
		}
		tasks[4] = tk
		wt.Done()
	}()*/
	go func() {
		tk := &TaskSpeed{
			Finished: 6,
			Total:    0,
		}
		tasks[5] = tk
		wt.Done()
	}()
	wt.Wait()

	for i := 0; i < 6; i++ {
		fmt.Printf("task_%d = %+v\n", i, tasks[i])
	}

}

/*实现一个安全的字符串*/
type SafeTask struct {
	missingTask string
	mx          sync.Mutex
}

func (s *SafeTask) deTask(rs string) () {
	s.mx.Lock()
	s.missingTask = strings.Replace(s.missingTask, rs, "", 1)
	s.mx.Unlock()
}

func main0() {
	var st = &SafeTask{
		missingTask: "123456",
		mx:          sync.Mutex{},
	}
	var wt = sync.WaitGroup{}
	wt.Add(5)
	go func(str string) {
		st.deTask(str)
		wt.Done()
	}("1")
	go func(str string) {
		st.deTask(str)
		wt.Done()
	}("2")
	go func(str string) {
		st.deTask(str)
		wt.Done()
	}("3")
	go func(str string) {
		st.deTask(str)
		wt.Done()
	}("4")
	go func(str string) {
		st.deTask(str)
		wt.Done()
	}("5")
	/*		go func(str string) {
			str = st.deTask(str)
			wt.Done()
		}("6")*/
	wt.Wait()

	fmt.Printf("s.missingTask = %s", st.missingTask)

}
