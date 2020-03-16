package main

import (
	"fmt"
	"regexp"
)
/*	bytes, e := ioutil.ReadFile("E:/workspace/Goland/正则/targetFile.txt")
	if e != nil {
		utils.HandleError(e,"ReadFile")
	}
	targetStr:=string(bytes)*/
func main() {
	targetStr := "a,aa,aaa,aaaa,aaaaa,aaaaaa"
	segStr1 := `a(a)*?`
	mustCompile := regexp.MustCompile(segStr1)
	submatch := mustCompile.FindAllString(targetStr, -1)
	fmt.Println(submatch)

}

//
func dome1() {
	targetStr := "我的电话:(17688700709)(13423059035),有时间给我电话！"
	segStr1 := `(\d{11}?)`
	mustCompile := regexp.MustCompile(segStr1)
	submatch := mustCompile.FindAllString(targetStr, -1)
	fmt.Println(submatch)
}

//()小括号切片输出需要获取值
func dome0()  {
	context2 := `
	        <title>标题</title>
	        <div>你过来啊</div>
	        <div>hello mike</div>
	        <div>你大爷</div>
	        <body>呵呵</body>
	    `
	segStr := "<div>(.*?)</div>"
	mustCompile := regexp.MustCompile(segStr)
	submatch := mustCompile.FindAllStringSubmatch(context2, -1)
	for _, v := range submatch {
		fmt.Println(v[1])
	}
}


