package main

import (
	"fmt"
	"strconv"
)

type sdxp struct{
	index int
	name string
}

type idriver interface{
	driver1()string
	driver2()string
}

func driver1()(string){
	fmt.Println("i am  driver1,i drive a car.")
	return "driver1"
}

func driver2()(string){
	fmt.Println("i am  driver2,i drive a traffic.")
	return "driver2"
}

func (s sdxp) driver1()(string){
	fmt.Println(s.name,"is a driver.")
	s.index = 3
	return "driver"+strconv.Itoa(s.index)
}

func (s *sdxp) driver2()(string){
	fmt.Println(s.name,"is a driver.")
	s.index = 5
	return "driver"+strconv.Itoa(s.index)
}

func Woker1(name string,num int){
	var sw sdxp
	var ssw *sdxp
	sw.index = num
	sw.name = name
	ssw = &sw
	driver1()
	driver2()
	fmt.Println("index is:",sw.index)
	sw.driver1()
	fmt.Println("index is:",sw.index)
	ssw.driver2()
	fmt.Println("index is:",sw.index)
}

func main() {
	Woker1("laowang",1)
}