package easyCode

import (
	"fmt"
	"testing"
)

func TestAddJsonFormGormTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
type CashOutBill struct {
BeHours string
ExamTime int
BulletinTime int
RestTime int
}
	`)
	fmt.Println(rs, len("QWERTYUIOPASDFGHJKLZXCVBNM1234567890"))
}
