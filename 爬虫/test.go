package main

import "fmt"

var chs = make(chan int, 100)
var ch00 = make(chan int, 1)

func main() {
	vv := 0
	ch00 <- vv
	inputchs(0)
	for v := range chs {
		fmt.Println(v, "\t", len(chs))

		go func(chs chan int) {
			fmt.Println("******",)
			if len(chs) == 5 {
				if vv < 6 {
					vv++
					ch00 <- vv
				}
			}
		}(chs)
	}

}

func inputchs(be int) {
	<-ch00
	en := be + 100
	for i := be; i < en; i++ {
		chs <- i
	}
}
