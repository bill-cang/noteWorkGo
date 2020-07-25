package utils

import (
	"math/rand"
	"strings"
	"time"
)

/*
最多保存35^35个数字，6为验证码最大为35^6-1
*/

var (
	InviteCode *inviteCode
)

const (
	ansInvCodeRand = "QWERTYUIOPASDGHJKLZXCVBNM1234567890" //邀请取码
)

type inviteCode struct {
	Base    string // 进制的包含字符, string类型
	Decimal uint64 // 进制长度
	Pad     string // 补位字符,若生成的code小于最小长度,则补位+随机字符, 补位字符不能在进制字符中
	Len     int    // code最小长度
}

func init() {
	//
	InviteCode = &inviteCode{
		Base:    ansInvCodeRand,
		Decimal: uint64(len(ansInvCodeRand)),
		Pad:     "F",
		Len:     6,
	}
}

// id转code
func (c *inviteCode) IdToCode(id uint64) string {
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % c.Decimal
		id = id / c.Decimal
		res += string(c.Base[mod])
	}

	resLen := len(res)
	if resLen < c.Len {
		res += c.Pad
		for i := 0; i < c.Len-resLen-1; i++ {
			rand.Seed(time.Now().UnixNano())
			res += string(c.Base[rand.Intn(int(c.Decimal))])
		}
	}
	return res
}

// code转id

func (c *inviteCode) CodeToId(code string) uint64 {
	res := uint64(0)
	lenCode := len(code)

	//var baseArr [] byte = []byte(c.Base)
	baseArr := []byte (c.Base)    // 字符串进制转换为byte数组
	baseRev := make(map[byte]int) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}

	// 查找补位字符的位置
	isPad := strings.Index(code, c.Pad)
	if isPad != -1 {
		lenCode = isPad
	}

	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == c.Pad {
			continue
		}
		index := baseRev[code[i]]
		b := uint64(1)
		for j := 0; j < r; j++ {
			b *= c.Decimal
		}
		res += uint64(index) * b
		r++
	}
	return res
}

func GetNewInviteCode(much int) *inviteCode {
	return &inviteCode{
		Base:    ansInvCodeRand,
		Decimal: uint64(len(ansInvCodeRand)),
		Pad:     "F",
		Len:     much,
	}
}
