package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
	"utils"
)

//获取所有老师的编号，并放入老师管道
func getHomePageArtists() {

	if StarPage == 1068 {
		close(chThid)
		return
	}

	for i := StarPage; i < EndPage; i++ {
		url := fmt.Sprintf("https://www.9fanhao.com/special-show-p-%d.htm", i)
		resp, err := http.Get(url)
		utils.HandleError(err, "http.Get")

		bytes, err := ioutil.ReadAll(resp.Body)
		utils.HandleError(err, "ioutil.ReadAll")
		pageStr := string(bytes)

		compile := regexp.MustCompile(repxYr)
		submatch := compile.FindAllStringSubmatch(pageStr, -1)

		fmt.Printf("Home第%d页，共%d人,缓存%d人\n", i, len(submatch), len(chThid))

		for _, val := range submatch {
			chThid <- val[1]
			//fmt.Print(val[1]+",")
		}

		StarPage++
		time.Sleep(4*time.Second)
		_ = resp.Body.Close()
	}

}
