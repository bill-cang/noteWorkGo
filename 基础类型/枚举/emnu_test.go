package eunm

import (
	"fmt"
	"testing"
)

func TestPolicyType_String(t *testing.T) {
	key := Policy_AVG.String()
	fmt.Println(key)
}
