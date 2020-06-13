package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddSensitiveToMap(t *testing.T) {
	words := strings.Split(InvalidWords, ",")
	for _, v := range words {
		InvalidWord[v] = nil
	}
	Set["你妈逼的"] = nil
	Set["你妈"] = nil
	Set["狗日"] = nil
	AddSensitiveToMap(Set)
	text := "文明用语你&* 妈, 逼的你这个狗 日的，怎么这么傻啊。我也是服了，我日,这些话我都说不出口"
	fmt.Println(ChangeSensitiveWords(text, sensitiveWord))
}

func TestChangeSensitiveWords(t *testing.T) {
	inviteCode := code{
		base:    "QWERTYUIOPASDGHJKLZXCVBNM1234567890",
		decimal: 35,
		pad:     "F",
		len:     8,
	}

	// 初始化检查
	if res, err := inviteCode.initCheck(); !res {
		fmt.Println(err)
		return
	}

	/*	code := inviteCode.idToCode(200219737)
		fmt.Printf("id=%v, code=%v\t", 200219737, code)
		//<-time.After(1 * time.Second)

		//mmp[code]=id
		//code = "HHC59YC8U6S"
		sid := inviteCode.codeToId(code)
		fmt.Printf("code=%v, id=%v\n", code, sid)*/

	mmp := make(map[string]int64)

	var id int64
	for id = 200219722; id < 200319722; id++ {
		ss := uint64(id)
		code := inviteCode.idToCode(ss)
		fmt.Printf("id=%v, code=%v\t", id, code)
		//<-time.After(1 * time.Second)

		//mmp[code]=id
		//code = "HHC59YC8U6S"
		sid := inviteCode.codeToId(code)
		fmt.Printf("code=%v, id=%v\n", code, sid)
	}

	fmt.Printf("ying : %d\nshi: %d", 200319722-200219722, len(mmp))
}
