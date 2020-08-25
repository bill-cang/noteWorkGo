package main

import "fmt"

func main() {
	a := 'a'
	for i := 1; i < 27; i++ {
		fmt.Printf("%d %c ", a, a)
		a++
	}
}

