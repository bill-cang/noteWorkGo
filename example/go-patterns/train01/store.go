package train01

import (
	"fmt"
	"noteWork/bestLib/utils/folatutil"
)

/*支付方式*/
type PAY_MODEL int

/*商店类型*/
type STORE_MODEL int

const (
	QIXISTORE STORE_MODEL = 1
	RTMARK    STORE_MODEL = 2
)

type Wares struct {
	UnitPrice float64 //单价
}

/*抽象商店行为*/
type Store interface {
	Sell(ws []string, model PAY_MODEL, payment float64) error //商品
	payModel
}

//商店基本属性
type StoreBase struct {
	PayDiscount map[PAY_MODEL]float64 //支付方式&折扣
	WaresMap    map[string]*Wares     //商品
	Turnover    float64               //营业额
	//*SerialOrder
}

//便利店实现商店行为
func (s *StoreBase) Sell(ws []string, sod *SerialOrder) (err error) {

	dsc, ok := s.PayDiscount[sod.PayMode]
	if !ok {
		return fmt.Errorf("No support payMode.")
	}

	total := float64(0)
	for _, v := range ws {
		ware, ok := s.WaresMap[v]
		if !ok {
			continue
		}
		total += ware.UnitPrice
	}
	total *= dsc
	total = folatutil.FloatUpwardMore(total, 2)

	if sod.Payment < total {
		return fmt.Errorf("Amount due %f.", total)
	}

	switch sod.PayMode {
	case PAY_CASH:
		sod.payCash(total)
	case PAY_ALI:
		sod.payAli(total)
	case PAY_WECHAT:
		sod.payWeChat(total)
	case PAY_CREDIT:
		sod.payCreDit(total)
	case PAY_HUAWEI:
		sod.payHuaWei(total)
	}

	s.Turnover += total
	return
}
