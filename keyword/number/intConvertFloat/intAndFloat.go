package infl

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func CalcMul() {
	decimal.DivisionPrecision = 2 // 保留两位小数，如有更多位，则进行四舍五入保留两位小数

	var num1 float64 = 3.36
	var num2 int = 2

	test := decimal.NewFromFloat(num1).Mul(decimal.NewFromFloat(float64(num2))).Round(0)
	f, _ := test.Float64()
	fmt.Println(int(f))
}
