package factory

import "fmt"

/*支付方式*/
type PAY_MODEL int

const (
	PAY_CASH   PAY_MODEL = 1
	PAY_ALI    PAY_MODEL = 2
	PAY_WECHAT PAY_MODEL = 3
	PAY_CREDIT PAY_MODEL = 4
	PAY_HUAWEI PAY_MODEL = 5
)

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

type payModel interface {
	payCash(float64) error
	payAli(float64) error
	payWeChat(float64) error
	payCreDit(float64) error
	payHuaWei(float64) error
}

type SerialOrder struct {
	PayMode  PAY_MODEL //支付方式
	Discount float64   //折扣
	Payment  float64   //收款
	Change   float64   //找零
}

/*七喜便利店*/
type QiXiStore struct {
	WaresMap map[string]*Wares //商品
	Turnover float64           //营业额
	*SerialOrder
}

func (s *SerialOrder) payCash(money float64) (err error) {
	s.Change = s.Payment - money*s.Discount
	return
}

func (s *SerialOrder) payAli(money float64) (err error) {
	s.Change = s.Payment - money*s.Discount
	return
}

func (s *SerialOrder) payWeChat(money float64) (err error) {
	s.Change = s.Payment - money*s.Discount
	return
}

func (s *SerialOrder) payCreDit(money float64) (err error) {
	s.Change = s.Payment - money*s.Discount
	return fmt.Errorf("Payment method not supported")
}

func (s *SerialOrder) payHuaWei(money float64) (err error) {
	s.Change = s.Payment - money*s.Discount
	return fmt.Errorf("Payment method not supported")
}

//便利店实现商店行为
func (x *QiXiStore) Sell(ws []string, model PAY_MODEL, payment float64) (err error) {
	total := float64(0)
	for _, v := range ws {
		ware, ok := x.WaresMap[v]
		if !ok {
			continue
		}
		total += ware.UnitPrice
	}

	if payment < total {
		return fmt.Errorf("Amount due %f.", total)
	}
	x.Payment = payment

	switch model {
	case PAY_CASH:
		err = x.payCash(total)
	case PAY_ALI:
		err = x.payAli(total)
	case PAY_WECHAT:
		err = x.payWeChat(total)
	case PAY_CREDIT:
		err = x.payCreDit(total)
	case PAY_HUAWEI:
		err = x.payHuaWei(total)
	}
	return
}

//大润发综超
type RTMark struct {
	Wares    []string //商品
	Turnover float64  //营业额
}
