package direc

import (
	"log"
	"os"
	"strings"
)

func GetObjectPath() (path string, err error) {
	//获取项目运行绝对路径
	var way string
	way, err = os.Getwd()
	if err != nil {
		log.Printf("[GetObjectPath] Getwd err =%+v", err)
		return
	}
	way = strings.Replace(way, "\\", "/", -1)
	return
}
