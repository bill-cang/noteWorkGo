package main

import "fmt"

var (
	res [100]int
)

func main() {

	for i := 0; i < 100; i++ {
		dome1(i)
	}

	fmt.Println(res)
}

func dome1(i int) {
	var tmp, flag = 0, 0
	if i%10 == 0 && flag == 0 {
		tmp = i
		error.Error("触发")
	}

	res[i]=i

	defer func(tmp int) {
		flag = 1
		go dome1(tmp)
	}(tmp)
}
