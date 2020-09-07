package train01

var QiXiStoreEnt *StoreBase

/*七喜便利店*/
func init() {
	initQiXiStore()
}

func initQiXiStore() {
	//初始化支持支付方式与折扣
	nmp := map[PAY_MODEL]float64{
		PAY_CASH:   1,
		PAY_ALI:    0.9,
		PAY_WECHAT: 0.95,
	}

	//初始化上架商品
	mmp := map[string]*Wares{
		"啤酒":  &Wares{UnitPrice: 5.0},
		"花生":  &Wares{UnitPrice: 8.5},
		"矿泉水": &Wares{UnitPrice: 2.5},
	}

	QiXiStoreEnt = &StoreBase{
		PayDiscount: nmp,
		WaresMap:    mmp,
		Turnover:    0,
	}

}
