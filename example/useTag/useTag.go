package useTag

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name        string `label:"name is: "`
	Age         int    `label:"age is: "`
	Gender      string `label:"gender is: " default:"unknown"`
}

func myPrint(obj interface{}) error {
	// 取 Value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)
	}

	return nil
}
