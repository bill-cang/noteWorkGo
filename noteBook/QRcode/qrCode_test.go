package QRcode

import (
	"fmt"
	"testing"
)

func TestGetQRCode(t *testing.T) {
	posterPath, _ := GetQRCode("E:/Tmp/最美校花/最美校花/0-开屏页1.png")
	fmt.Println(posterPath)
}
