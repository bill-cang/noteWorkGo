package folatutil

import (
	"fmt"
	"math"
	"strings"
)

//保留小数点后n位，只考虑对n+1位向上保留
func FloatUpwardLess(fo float64, l uint8) float64 {
	d := math.Pow10(int(l))
	return math.Trunc(fo*d) / d
}

//保留小数点后n位，对n位后存在有效位向上保留
func FloatUpwardMore(fo float64, l uint8) float64 {
	sf := fmt.Sprintf("%g", fo)
	st := strings.Split(sf, ".")
	if len(st[1]) > int(l) {
		//存在有效位，为保留位
		fo = fo + math.Pow10(-int(l))
	}
	return FloatUpwardLess(fo, l)
}

//保留小数点后n位，对n位向下保留
func FloatDownward(fo float64, l uint8) float64 {
	d := math.Pow10(int(l))
	return math.Trunc(fo*d) / d
}

func getAddend(l uint8) float64 {
	switch l {
	case 1:
		return 1e-1
	case 2:
		return 1e-2
	case 3:
		return 1e-3
	case 4:
		return 1e-4
	case 5:
		return 1e-5
	case 6:
		return 1e-6
	case 7:
		return 1e-7
	case 8:
		return 1e-8
	case 9:
		return 1e-9
	}
	return 1e-1
}
