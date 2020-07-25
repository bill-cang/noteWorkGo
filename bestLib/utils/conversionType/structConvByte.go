package conversionType

import (
	"log"
	"noteWork/bestLib/terror"
	"reflect"
	"unsafe"
)

//https://www.jb51.net/article/144265.htm

type sliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func StructConversionByte(inf interface{}) (bytes []byte, err error) {
	kind := reflect.TypeOf(inf).Kind()
	switch kind {
	case reflect.Struct:
		bytes = []byte{}
		Len := unsafe.Sizeof(inf)
		tempBytes := &sliceMock{
			addr: uintptr(unsafe.Pointer(&inf)),
			cap:  int(Len),
			len:  int(Len),
		}
		bytes = *(*[]byte)(unsafe.Pointer(tempBytes))
	default:
		log.Printf("[StructConversionByte] err inf type is %+v\n", kind)
		err = terror.ErrBadParam
	}
	return
}

func BytesConversionStruct(bytes *[]byte, tp reflect.Type) (inf interface{}) {
	//ptestStruct *TestStructTobytes = *(**TestStructTobytes)(unsafe.Pointer(&data))

	/*	if tp.Kind() != reflect.Struct {
			return
		}
		inf = (**tp)(unsafe.Pointer(bytes))*/
	return
}
