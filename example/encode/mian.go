package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
)

//md5值
func Md5Str(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}


// 散列值
func Sha1Str(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func HashCode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func main() {
	txt := "https://github.com/hashicorp/terraform/blob/master/helper/hashcode/hashcode.go"
	fmt.Println(HashCode(txt))
	fmt.Println(Md5Str(txt))
	fmt.Println(Sha1Str(txt))
}
