package m_struct

import (
	"errors"
	"reflect"
)

type Animal struct {
	subject string `json:"subject"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
}

type Dog struct {
	Animal
	Trait string
}

type Cat struct {
	Name string
	Age  int
}

func (s *Cat)ToMap() (mmp map[string]interface{},err error)  {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(e.(string))
			return
		}
	}()
	mmp = make(map[string]interface{})
	obj1 := reflect.TypeOf(*s)
	obj2 := reflect.ValueOf(*s)
	for i := 0; i < obj1.NumField(); i++ {
		mmp[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return
}
