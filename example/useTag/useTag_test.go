package useTag

import (
	"encoding/json"
	"fmt"
	"noteWork/bestLib/utils/encode"
	"testing"
)

func TestPrint(t *testing.T) {

	type AccountBody struct {
		Name     string
		Password string
		Id       string
		Words    []string
		KeyStore string
	}

	ab := &AccountBody{
		Words: []string{
			"wreck", "fame",
			"update", "void", "couch",
		},
	}

	//det := []byte("0709")
	json.Marshal(ab)

	pk := "ckx070900000"
	aes := encode.EncryptAES([]byte("123456789"), []byte(pk))
	decryptAES := encode.DecryptAES(aes, []byte(pk))

	/*	encode := hex.EncodeToString(bytes)
		decodeString, err := hex.DecodeString(encode)*/
	fmt.Printf("aes = %s\ndecryptAES =%s\n", string(aes), string(decryptAES))
}



