package m_struct

import (
	"fmt"
	"testing"
)

func TestStructToMap(t *testing.T) {
	var cat = Cat{
		Name: "ckx",
	}

/*	var cat = Animal{
		subject: "123",
		Name:    "4567",
		Age:     0,
	}*/
	fmt.Printf("tomap = %s%%\n", "ckx")

	toMap, err := cat.ToMap()
	if nil != err {
		fmt.Printf("err = %+v\n", err)
	}
	fmt.Printf("tomap = %+v\n", toMap)
}
