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
