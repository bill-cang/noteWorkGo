package factory

import (
	"fmt"
	"testing"
)

func TestGeneratePayment3(t *testing.T) {
	var store Store
	mmp := map[string]*Wares{
		"啤酒":  &Wares{UnitPrice: 5.0},
		"花生":  &Wares{UnitPrice: 8.5},
		"矿泉水": &Wares{UnitPrice: 2.5},
	}

	store = &QiXiStore{
		WaresMap:    mmp,
		Turnover:    0,
		SerialOrder: &SerialOrder{},
	}

	err := store.Sell([]string{"啤酒", "花生"}, PAY_ALI, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	x := store.(*QiXiStore)
	fmt.Printf("收款：%f, 找零：%f, 营业额：%f", x.Payment, x.Change, x.Turnover)
}
