package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with  %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()

	runTimes := 1000

	// Use the common pool.
	syncCalculateSum := func() {
		demoFunc()
	}
	for i := 0; i < runTimes; i++ {
		_ = ants.Submit(syncCalculateSum)
	}
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function, 
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{})  {
		myFunc(i)
	},)
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		_ = p.Invoke(int32(i))
	}
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)

	<-time.After(time.Second * 6)
}
