package main

import (
	"fmt"
	"github.com/robfig/cron"
	"noteWork/keyword/time/timeTools"
)

/*教程：
https://www.jianshu.com/p/e629d637bf4c
*/

type TestJob struct {
}

func (this TestJob)Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job)Run() {
	fmt.Println("testJob2...")
}

//启动多个任务
func main() {
	i := 0
	c := cron.NewWithLocation(timeTools.GetTimeLocation("KE"))

	//AddFunc
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})

	//AddJob方法
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})

	//启动计划任务
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select{}
}