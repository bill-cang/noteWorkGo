package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//打开数据库
	openDBconnect()
}

func main() {
	chHome <- 1

	//分析home页数据，获取所有老师id，并放入缓存管道
	go getHomePageArtists()

	go getPageDataToChan()

	//开辟10条协程固化老师信息到数据库
	for i := 0; i < 10; i++ {
		go insertArtist()
	}

	//开辟10条协程固化作品信息到本地数据库
	for i := 0; i < 100; i++ {
		go insertVideos()
	}

	//开辟10条协程固化作品信息到本地数据库
	for i := 0; i < 10; i++ {
		go insertFail()
	}

	//开辟20条协程分别请求20位老师地址
/*	for i := 0; i < 100; i++ {
		go getPageDataToChan()
	}*/

	/*wg.Add(1)
	go func() {
		for {
			time.Sleep(30 * time.Second)
			//fmt.Println("chVido-》",len(chVido),"chVdoCount：",len(chVdoCount),"chVdoHandle",len(chVdoHandle))
			//if len(chVdoCount) == len(chVdoHandle) && StarPage==1059{
			fmt.Printf("GetPage=%d,chThEt=%d,chVido=%d,chFail=%d",
				GetPage, len(chThEt), len(chVido), len(chFail))
			if GetPage == 2 && len(chVido) == 0 && len(chFail) == 0 {
				close(chThEt)
				close(chVido)
				//time.Sleep(30 * time.Second)
				close(chFail)
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()*/
	over := <-chExit
	defer func() {
		close(chThEt)
		close(chVido)
		close(chFail)
	}()
	fmt.Println("$$$数据爬取完毕 %d", over)
}
