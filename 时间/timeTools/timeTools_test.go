package timeTools

import (
	"fmt"
	"testing"
)

func TestGetAfterSomeDateOriginStamp(t *testing.T) {
	str := GetRealTimeStr(Tf_complete, TimeLayout)
	fmt.Println(str)
	time, _ := GetTimeFromDateTime("1582101125")
	fmt.Println(time)
}
