package main

import (
	"errors"
	"fmt"
)

func main() {

	slc := []int{1, 2, 3, 4, 5}
	slc2 := make([]int, 4)
	copy(slc2, slc)
	slc2[0] = 9
	slc2 = append(slc2, 8)
	fmt.Printf("slc =%+v\nslc2 =%+v", slc, slc2)

	test1()

}

func test1() (err error) {
	fmt.Printf("Begin err p =%p\n", err)
	if age, err := test2(); err != nil {
		fmt.Printf("age = %d,err p = %p\n", age, err)
		return err
	}
	return
}

func test2() (age int, err error) {
	return 1, errors.New("There is something.")
}
