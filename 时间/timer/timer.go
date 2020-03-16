package timer

import (
	"fmt"
	"time"
)

func NextZeroClockTimer() {
	for {
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, util.GetTimeLocation("Kenya"))
		t := time.NewTimer(next.Sub(now))
		<-t.C

		doSomething()

	}
}

func doSomething() {
	fmt.Println("零时定时任务！")
}
