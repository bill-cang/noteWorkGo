package folatutil

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	fo0 = 1.2391
	fo1 = 1.2301
	fo5 = 1.2305
	fo9 = 1.2399
)

func TestFloatKeepDot(t *testing.T) {
	dot1 := FloatUpwardLess(fo1, 2)
	dot2 := FloatUpwardMore(fo1, 2)
	fmt.Printf("dot1 =%g\tdot2 =%g\n", dot1, dot2)
}

func TestFloatDownwardLess(t *testing.T) {
	less := FloatDownward(fo9, 2)
	fmt.Print(less)
}

func TestFloatDownward(t *testing.T) {
	//f := math.Pow10(-2)
	/*	f := 1e-2
		fmt.Printf("f =%f\n", f)*/

}

func TestFloatUpwardLess(t *testing.T) {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", 19.90), 64)
	fmt.Println(num)
}

func TestFloatUpwardMore(t *testing.T) {
	v1, v2 := "1", "2"
	fmt.Print(v1 + v2)
}
