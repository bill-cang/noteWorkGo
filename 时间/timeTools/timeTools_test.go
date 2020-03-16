package timeTools

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestGetAfterSomeDateOriginStamp(t *testing.T) {

	now := time.Now().Unix()
	formatStr := GetFormatDateFormLocation("", 1583251200)
	fmt.Println("now = ", now, "formatStr = ", formatStr)
	//1583251373
	stamp, _ := GetDateStrToStamp("", "2020-03-04 00:02:53")
	fmt.Println("stamp = ", stamp)

	zeroStamp := GetAfterSomeDateZeroClockStamp("", 0)
	fmt.Printf("zeroStamp =%d", zeroStamp)

	str := "123456"
	str = strings.Replace(str, "1", "", 1)
	str = strings.Replace(str, "2", "", 1)
	str = strings.Replace(str, "3", "", 1)
	str = strings.Replace(str, "4", "", 1)
	str = strings.Replace(str, "5", "", 1)
	str = strings.Replace(str, "6", "", 1)

	fmt.Println(str)

}

func TestGetEqualDay(t *testing.T) {
	eq := GetEqualDay(86400)
	fmt.Printf("eq := %+v\n", eq)
}
