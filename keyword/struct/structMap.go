package m_struct

import (
	"errors"
	"reflect"
)

//反射不能获取结构体的私有属性 ：reflect.Value.Interface: cannot return value obtained from unexported field or method
func StructToMap(obj interface{}) (mmp map[string]interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(e.(string))
			return
		}
	}()
	mmp = make(map[string]interface{})
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	for i := 0; i < obj1.NumField(); i++ {
		mmp[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return
}

