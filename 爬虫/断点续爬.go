package main
//断电续爬思路
import (
	"fmt"
	"sync"
	"time"
)

var (
	ch001  chan int
	ch002  chan int
	wg     sync.WaitGroup
	cunMap = make(map[int]string, 100)
)

func init() {
	ch001 = make(chan int, 1000)
	ch002 = make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		ch001 <- i
	}
}

func main() {

	wg.Add(1)
	go func() {
		for {
			fmt.Println("######", len(ch002))
			if len(ch002) == 1000 {
				fmt.Println("######!!!", len(ch002))
				close(ch001)
				break
			}
		}
		close(ch002)
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		go dom001()
	}

	wg.Wait()
}

func dom001() {
	var errInt int

	defer func() {
		if err := recover(); err != nil {
			errStr := fmt.Sprintf("%s", err)
			if errStr == "runtime error: integer divide by zero" {
				fmt.Println("******", errInt)
				cunMap[errInt] = "hulue"
				time.Sleep(3 * time.Second)
				dom001()
				ch001 <- errInt
			}
		}
	}()
	for val := range ch001 {
		q := 1
		errInt = val
		fmt.Println("***", val)
		cun := cunMap[val]
		if val%100 == 0 && val != 0 && cun != "" {
			q = 0
		}

		shang := errInt / q
		fmt.Println(shang)
		ch002 <- shang
	}
}
