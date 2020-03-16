package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"utils"
)

var (
	//<a href="/special-read-id-11978.html" target="_blank">
	//repxYr = `<a href="/special-read-id-\d{1,5}.html" target="_blank"><img`
	repxYr=`href="/special-read-id-(\d{1,5}).html" target="_blank"><img`
)

func main() {
	url := fmt.Sprintf("https://www.9fanhao.com/special-show-p-%d.htm", 10)
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	utils.HandleError(err, "ioutil.ReadAll")
	pageStr := string(bytes)
	fmt.Println(pageStr)

	compile := regexp.MustCompile(repxYr)
	submatch := compile.FindAllStringSubmatch(pageStr, -1)

	fmt.Printf("第%d页，共%d人\n", 10, len(submatch))
	for _, val := range submatch {
		fmt.Println(val[0])
	}
}
