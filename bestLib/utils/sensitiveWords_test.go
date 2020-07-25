package utils

import (
	"fmt"
	"noteWork/bestLib/utils/conversionType"
	"reflect"
	"testing"
)

func TestAddSensitiveToMap(t *testing.T) {
	text := "文明用语你&* 妈, 逼的你这个狗 日的，怎么这么傻啊。我也是服了，我日,这些话我都说不出口bo1"
	fmt.Println(ChangeSensitiveWords(text, sensitiveWord))
}

func BenchmarkChangeSensitiveWords(b *testing.B) {
	sl1 := []string{"a", "v", "g", "r", "l", "d", "f", "t", "r"}
	sl2 := []string{"f", "e", "6", "&", "i", "p", "l", "v", "t"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SlcIntersect(sl1, sl2)
	}
}

func TestChangeSensitiveWords(t *testing.T) {
	type sliceMock struct {
		addr uintptr
		len  int
		cap  int
	}

	type animal struct {
		name string
		age  int
	}

	var testStruct = animal{
		name: "zk",
		age:  0,
	}

	bytes, err := conversionType.StructConversionByte(testStruct)
	typeOf := reflect.TypeOf(testStruct)
	fmt.Printf("bytes =%+v, err =%+v ,typeOf =%+v", bytes, err, typeOf)

}
