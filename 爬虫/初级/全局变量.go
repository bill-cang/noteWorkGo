package main

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

var (
	//chdata = make(chan string, 10000)
	chVido = make(chan *video, 1000)   //作品管道
	chThid = make(chan string, 1000)   //主页老师id管道
	chThEt = make(chan *teacher, 1000) //老师实体管道
	chFail = make(chan *fails, 1000)
	DB     *sqlx.DB
	wg     sync.WaitGroup

	StarPage =1067//home页开始爬虫页码
	EndPage=1068
	//GetPage  =1

	chHome = make(chan int,1)
	chExit = make(chan int,1)

	/*	thEtCount   = make(chan int, 10e4)
		thEtHandle = make(chan int, 10e4)*/

/*	chVdoCount  = make(chan int, 10e4) //统计写入管道的总影片数
	chVdoHandle = make(chan int, 10e4)*/ //统计写入数据库的总影片数

	err001 = "connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond"
)
