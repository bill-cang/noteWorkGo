package custonErr

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	err := ErrBadParamsBody
	fmt.Printf("err =%+v\n", err)
}
