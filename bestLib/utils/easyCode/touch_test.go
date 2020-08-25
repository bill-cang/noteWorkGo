package easyCode

import (
	"fmt"
	"testing"
)

func TestAddJsonFormGormTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
type NearbyProfile struct {
	Id             int64
	Uid            string
	Interests      string
	Personality    string
	Looks          string
	Height         int32
	Mode           int32
	FilterSex      int32
	FilterAge      string
	FilterDistance int32
	Hide           int32
	Identity       int32
	CreateTime     int64
}
	`)
	fmt.Println(rs, len("QWERTYUIOPASDFGHJKLZXCVBNM1234567890"))
}


