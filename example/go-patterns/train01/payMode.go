package train01

const (
	PAY_CASH   PAY_MODEL = 1
	PAY_ALI    PAY_MODEL = 2
	PAY_WECHAT PAY_MODEL = 3
	PAY_CREDIT PAY_MODEL = 4
	PAY_HUAWEI PAY_MODEL = 5
)

type payModel interface {
	payCash(float64)
	payAli(float64)
	payWeChat(float64)
	payCreDit(float64)
	payHuaWei(float64)
}

//实现支付接口
type SerialOrder struct {
	PayMode PAY_MODEL //支付方式
	//Discount float64   //折扣
	Payment float64 //收款
	Change  float64 //找零
}

func (s *SerialOrder) payCash(money float64) {
	s.Change = s.Payment - money
	return
}

func (s *SerialOrder) payAli(money float64) {
	s.Change = s.Payment - money
	return
}

func (s *SerialOrder) payWeChat(money float64) {
	s.Change = s.Payment - money
	return
}

func (s *SerialOrder) payCreDit(money float64) {
	s.Change = s.Payment - money
	return
}

func (s *SerialOrder) payHuaWei(money float64) {
	s.Change = s.Payment - money
	return
}
