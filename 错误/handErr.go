package errs

import (
	"fmt"
	"runtime"
	"strings"
)

func HandleError(err error, where string) {
	if err==nil{
		return
	}
	errStr := fmt.Sprintf("%s", err)
	fmt.Printf("■■■%s[%s]  -------->>%s\n",printMyName(),where, errStr)
}

//获取调用函数的funcName,
func printMyName() string {
	//skip ,可以调节获取协程
	pc, _, _, _ := runtime.Caller(2)
	funNmae:=runtime.FuncForPC(pc).Name()
	index := strings.LastIndex(funNmae, "/")+1
	return funNmae[index:]
}