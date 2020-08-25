package ChinaGo

import "fmt"

/*
1. iota 是Go语言的常量计数器，只能在常量表达式中使用
2. iota 在const关键字出现的位置将被重置为0，const中每一行被计数一次
*/
const (
	a = iota
	b = iota
)

const (
	name = "name"
	c    = iota
	d    = iota
	age  = 30
	e    = iota
)
//com 01
//com 02

func DOPrint() {
	fmt.Printf("a = %d\nb =%d\nc =%d\nd = %d\ne = %d\n", a, b, c, d, e)
}
