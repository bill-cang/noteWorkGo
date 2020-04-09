package timeTools

import (
	"fmt"
	"testing"
	"time"
)

func TestGetAfterSomeDateOriginStamp(t *testing.T) {

	now := time.Now().Unix()

	//1583251373
	stamp, _ := GetDateStrToStamp("", "2020-03-28 10:44:00")
	fmt.Println("stamp = ", stamp)

	formatStr := GetFormatDateFormLocation("", stamp)
	fmt.Println("now = ", now, "formatStr = ", formatStr)

}

func TestGetEqualDay(t *testing.T) {
	formatDate := GetFormatDateFormLocation("Kenya", 1585542780)
	fmt.Printf("formatDate :%s\n",formatDate)

	//time.Unix(1585224483,0).Format("01/02/2006 15:04:05")
	simpleTimeFormat := time.Unix(1585224483,0).In(GetTimeLocation("Kenya")).Format("01/02/2006")
	fmt.Printf("simpleTimeFormat =%s\n",simpleTimeFormat)
}



