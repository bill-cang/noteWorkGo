package encryption

import (
	"fmt"
	"noteWork/utils"
	"testing"
)

func TestMd5SumStr(t *testing.T) {
	str := "hello world"
	digest, err := Md5SumStr(str)
	if err != nil || len(digest) != 32 {
		t.Fail()
	}
}

func TestMd5SumRaw(t *testing.T) {
	str := "hello world"
	digest, err := Md5SumRaw(str)
	if err != nil || len(digest) != 16 {
		t.Fail()
	}
}

func TestMd5File(t *testing.T) {
	file := "E:/Tmp/Greetings audio aac/女声/Essy.aac"
	digest, err := Md5File(file)
	if err != nil || len(digest) != 32 {
		t.Fail()
	}
	size := utils.GetFileSize(file)
	fmt.Printf("digest = %s , Size = %d\n", digest, size)
}
