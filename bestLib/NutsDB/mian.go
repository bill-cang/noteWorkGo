package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomString, _ := GetRandomString(10, 'a','b','f')

	fmt.Print(randomString)

}

// GetRandomString generate random string by specify chars.
// 通过指定字符生成随机字符串
func GetRandomString(n int, alphabets ...byte) (string, error) {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		if len(alphabets) == 0 {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes), nil
}
