package main

import (
	"io/ioutil"
	"net/http"
	"utils"
)

func getHtmlStr() *string{
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	//关闭资源
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	htmlStr := string(bytes)
	//fmt.Println(htmlStr)
	return &htmlStr
}
