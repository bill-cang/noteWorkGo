package tools

import (
	"errors"
	"fmt"
	"reflect"
)

func IsInSlice(sc []interface{}, element interface{}) {

}

/*==============================>切片去重*/
/* 在slice中去除重复的元素，其中a必须是已经排序的序列。
* params:
* a: slice对象，如[]string, []int, []float64, ...
* return:
* []interface{}: 已经去除重复元素的新的slice对象
 */
func SliceRemoveDuplicate(a interface{}) (ret []interface{}, err error) {
	if reflect.TypeOf(a).Kind() != reflect.Slice {
		fmt.Printf("<SliceRemoveDuplicate> <a> is not slice but %T\n", a)
		err = errors.New(fmt.Sprintf("a.type =%t,not splic type.", reflect.TypeOf(a)))
		return
	}
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return
}
/*==============================>切片插入*/
/*
* 在Slice指定位置插入元素。
* params:
* s: slice对象，类型为[]interface{}
* index: 要插入元素的位置索引
* value: 要插入的元素
* return:
* 已经插入元素的slice，类型为[]interface{}
 */

func SliceInsert(s []interface{}, index int, value interface{}) []interface{} {
	rear := append([]interface{}{}, s[index:]...)
	return append(append(s[:index], value), rear...)
}

/*
* 在Slice指定位置插入元素。
* params:
* s: slice对象指针，类型为*[]interface{}
* index: 要插入元素的位置索引
* value: 要插入的元素
* return:
* 无
 */

func SliceInsert2(s *[]interface{}, index int, value interface{}) {
	rear := append([]interface{}{}, (*s)[index:]...)
	*s = append(append((*s)[:index], value), rear...)
}


/*
* 在Slice指定位置插入元素。
* params:
* s: slice对象的指针，如*[]string, *[]int, ...
* index: 要插入元素的位置索引
* value: 要插入的元素
* return:
* true: 插入成功
* false: 插入失败（不支持的数据类型）
 */

func SliceInsert3(s interface{}, index int, value interface{}) bool {
	if ps, ok := s.(*[]string); ok {
		if val, ok := value.(string); ok {
			rear := append([]string{}, (*ps)[index:]...)
			*ps = append(append((*ps)[:index], val), rear...)
			return true
		}
	} else if ps, ok := s.(*[]int); ok {
		if val, ok := value.(int); ok {
			rear := append([]int{}, (*ps)[index:]...)
			*ps = append(append((*ps)[:index], val), rear...)
		}
	} else if ps, ok := s.(*[]float64); ok {
		if val, ok := value.(float64); ok {
			rear := append([]float64{}, (*ps)[index:]...)
			*ps = append(append((*ps)[:index], val), rear...)
		}
	} else {
		fmt.Printf("<SliceInsert3> Unsupported type: %T\n", s)
	}
	return false
}


/*==============================>切片插入*/
