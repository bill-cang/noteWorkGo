package encryption

import (
	"encoding/json"
	"fmt"
	"noteWork/bestLib/utils"
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
	/*
		"robot","pave","cloth",
		"exist","vast","carpet",
	*/

	/*str := "hello world"
	digest, err := Md5SumRaw(str)
	if err != nil || len(digest) != 16 {
		t.Fail()
	}*/

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

type animal struct {
	name string
	age  int
}

func TestMd5Sum(t *testing.T) {

	var ani = &animal{
		name: "zk",
		age:  0,
	}
	var ani1 = &animal{
		name: "zk",
		age:  23,
	}

	//str, err := Md5SumStr(ani)
	bytes, _ := json.Marshal(ani)
	st1 := &animal{}
	json.Unmarshal(bytes, st1)
	//fmt.Printf("str = %s, err = %+v , bytes = %+v\n", str, err, bytes)
	fmt.Printf("bytes =%+v , st =%+v\n", bytes, st1)

	//str1, err := Md5SumStr(ani1)
	bytes2, _ := json.Marshal(ani1)
	st2 := &animal{}
	json.Unmarshal(bytes2, st2)
	//fmt.Printf("str= %s, err = %+v , bytes = %+v\n", str1, err, bytes)
	fmt.Printf("bytes =%+v , st =%+v\n", bytes2, st2)

	//[Md5SumStr] data = &[123 125]str= 99914b932bd37a50b983c5e7c90ae93b, err = <nil> , bytes = [34 57 57 57 49 52 98 57 51 50 98 100 51 55 97 53 48 98 57 56 51 99 53 101 55 99 57 48 97 101 57 51 98 34]

}
