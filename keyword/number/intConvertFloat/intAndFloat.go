package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	price, err := decimal.NewFromString("1.2345678")
	if err != nil {
		panic(err)
	}
	//保留小数位，四舍五入
	rd := price.Round(3)
	//保留小数位，向上取
	cash := price.RoundCash(5)

	f, of := rd.Float64()
	if of {
		return
	}
	fmt.Printf("mul :=%f ,cash =%+v", f, cash)
}

func main1() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
