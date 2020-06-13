package main

import "fmt"

func main() {
	a := make(chan bool, 100)
	b := make(chan bool, 100)
	c := make(chan bool, 100)
	for i := 0; i < 10; i++ {
		a <- true
		b <- true
		c <- true
	}
	for i := 0; i < 10; i++ {
		select {
		case <-a:
			fmt.Printf("< a =%v\n",a)

		case <-b:
			fmt.Printf("< b =%v\n",b)

		case <-c:
			fmt.Printf("< c =%v\n",c)

		default:
			fmt.Printf("< default\n")
		}
	}
}
