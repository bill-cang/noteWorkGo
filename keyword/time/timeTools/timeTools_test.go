package timeTools

import (
	"fmt"
	"testing"
	"time"
)

func TestGetAfterSomeDateOriginStamp(t *testing.T) {

	now := time.Now().Unix()
	formatStr := GetFormatDateFormLocation("", now)
	fmt.Println("now = ", now, "formatStr = ", formatStr)

	//1583251373 20200610200000
	stamp, _ := GetDateStrToStamp("", "2020-05-26 00:12:00")
	fmt.Println("stamp = ", stamp)
}

func TestGetEqualDay(t *testing.T) {
	formatDate := GetFormatDateFormLocation("Kenya", 1591635600)
	fmt.Printf("formatDate :%s\n",formatDate)

	//time.Unix(1585224483,0).Format("01/02/2006 15:04:05")
	simpleTimeFormat := time.Unix(1587208207,0).In(GetTimeLocation("Kenya")).Format("01/02/2006")
	fmt.Printf("simpleTimeFormat =%s\n",simpleTimeFormat)
}



