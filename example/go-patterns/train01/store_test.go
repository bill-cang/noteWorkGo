package train01

import (
	"fmt"
	"github.com/willf/bitset"
	"testing"
)

func TestGeneratePayment3(t *testing.T) {

	sod := &SerialOrder{
		PayMode: PAY_WECHAT,
		Payment: 20,
	}

	err := QiXiStoreEnt.Sell([]string{"啤酒", "花生"}, sod)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("收款：%f, 找零：%f, 营业额：%f\n", sod.Payment, sod.Change, QiXiStoreEnt.Turnover)

}

func TestInitQiXiStoreEnt(t *testing.T) {
	var bst bitset.BitSet
	bst.Set(9)
	fmt.Printf("bst len =%d , bool9 = %t", bst.Len(), bst.Test(9))
}

func TestStoreBase_Sell(t *testing.T) {

	/*	sprintf := fmt.Sprintf("%.3f", 1.2345)
		fmt.Println(sprintf)*/

	/*	float := folatutil.FormatFloat(1.2345, 2)
		fmt.Printf("%s\n", float)*/

	/*	trunc := math.Trunc(1.2345)
		fmt.Print(trunc)*/

	/*	d := 1.20/1.2
		fmt.Print(d)*/

	//fmt.Print(math.Trunc(1.2395*1e2+0.6) * 1e-2)



	/*	nf := fmt.Sprintf("%g", 1.23400)
		fmt.Print(nf)*/

}
