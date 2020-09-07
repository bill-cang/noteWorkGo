package factory

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	balance float32 = 100.00
)

func TestGeneratePayment(t *testing.T) {
	for i := 0; i < 4; i++ {
		if payment, err := GeneratePayment(Kind(i), balance); err != nil {
			fmt.Printf("GeneratePayment Kind =%d , err =%+v\n", i, err)
			continue
		} else {
			fmt.Println(reflect.TypeOf(payment).Elem().String(), "paling....")
			if err := payment.Pay(38.7); err != nil {
				fmt.Printf("GeneratePayment Kind =%d , err =%+v\n", i, err)
			}

			switch Kind(i) {
			case Cash:
				fmt.Printf("paly success. Balance =%f\n", payment.(*CashPay).Balance)
			case Credit:
				fmt.Printf("paly success. Balance =%f\n", payment.(*CashPay).Balance)
			}
		}
	}
}

func TestCashPay_Pay(t *testing.T) {
	payment, _ := GeneratePayment(1, balance)
	payment.Pay(20)
	//cash := reflect.New(reflect.TypeOf(payment).Elem()).Interface().(*CashPay)relect新的对象
	cash := payment.(*CashPay)
	fmt.Println(reflect.TypeOf(cash))
	if cash.Balance != float32(80) {
		t.Error("结算错误")
	}
}

func TestGeneratePayment2(t *testing.T) {
	fmt.Printf("Cash =%d,Credit =%d", Cash, Credit)
}
