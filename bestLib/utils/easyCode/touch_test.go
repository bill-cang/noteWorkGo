package easyCode

import (
	"fmt"
	"testing"
)

func TestAddJsonFormGormTag(t *testing.T) {
	rs := AddJsonFormGormTag(`
type notice struct {
	id          int32
	uid         string
	ntc_type    int32
	ntc_body    string
	trigger_uid string
	create_time int64
}
	`)
	fmt.Println(rs, len("QWERTYUIOPASDFGHJKLZXCVBNM1234567890"))
}


