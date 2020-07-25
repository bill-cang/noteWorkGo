package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"noteWork/bestLib/utils/conversionType"
	"noteWork/bestLib/utils/zerocopy"
	"os"
	"reflect"
	"strings"
)

func Md5SumStr(iData interface{}) (digest string, err error) {
	var data *[]byte
	kind := reflect.TypeOf(iData).Kind()
	switch kind {
	case reflect.String:
		tmp := zerocopy.StrToBytes(iData.(string))
		data = &tmp
	case reflect.Struct:
		/*		if data, err = conversionType.StructConversionByte(iData); err != nil {
				return
			}*/
		if tmp, err := json.Marshal(iData); err != nil {
			log.Printf("[Md5SumStr] json.Marshal iData =%+v", iData)
			return "" ,err
		} else {
			data = &tmp
		}

		fmt.Printf("[Md5SumStr] data = %+v", data)
	default:
		err = errors.New("Unknown iData type")
		return
	}
	h := md5.Sum(*data)
	digest = hex.EncodeToString(h[:])
	return
}

func hEncode(bs *[16]byte) string {
	return hex.EncodeToString(bs[:])
}

func Md5SumRaw(iData interface{}) (digest []byte, err error) {
	var data []byte
	switch vv := iData.(type) {
	case string:
		data = zerocopy.StrToBytes(vv)
	case []byte:
		data = vv
	case struct{}:
		if data, err = conversionType.StructConversionByte(iData); err != nil {
			return
		}
	default:
		err = errors.New("buxiaode")
		return
	}
	h := md5.Sum(data)
	digest = h[:]
	return
}

func Md5File(file string) (digest string, err error) {
	var fp *os.File
	fp, err = os.Open(file)
	if err != nil {
		return
	}
	defer fp.Close()
	m := md5.New()
	io.Copy(m, fp)
	h := m.Sum(nil)
	digest = hex.EncodeToString(h)
	return
}

// Md5Sum calculates the md5sum of a stream
// 计算流的md5sum
func Md5Sum(reader io.Reader) (string, error) {
	var returnMD5String string
	hash := md5.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

// Md5SumString calculates the md5sum of a string
// 计算字符串的md5sum
func Md5SumString(input string) (string, error) {
	buffer := strings.NewReader(input)
	return Md5Sum(buffer)
}
