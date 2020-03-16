package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"utils"
)

//建立数据库链接
func openDBconnect() {
	//★★★★★曾经有这么一个大坑！！！全局变量DB因为和局部变量err一起使用“:=”,DB也变成了局部变量！！造成了其他地方的空指针！！
	var err error
	DB, err = sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/actress")
	DB.SetMaxOpenConns(2000)
	DB.SetMaxIdleConns(1000)
	if DB != nil {
		fmt.Println("==数据库链接成功！", DB)
	} else {
		fmt.Println("==数据库链接失败！", DB)
		utils.HandleError(err, "sqlx.Connect")
	}
}

//将老师们的作品信息存入数据库 issuedate
func insertVideos() {
	sql := "INSERT INTO VIDEOS (CODE,TITLE,ACTRESS,ISSUEDATE,FILMLENGTH) VALUES (?,?,?,?,?);"
	for ap := range chVido {
		fmt.Println("chVido->", len(chVido), "：", ap.code, "\t", ap.title, "\t", ap.teacher, "\t", ap.filmlength, "\t", ap.issueData)
		code, title, actress, issuedate, filmlength := ap.code, ap.title, ap.teacher, ap.issueData, ap.filmlength
		_, e := DB.Exec(sql, code, title, actress, issuedate, filmlength)
		utils.HandleError(e, "DB.Exec VIDEOS")
		//chVdoHandle<-1
	}
}

//接收老师们的个人信息，固化到数据库
func insertArtist() {
	sql := "INSERT INTO `artist`( `name`, `ename`, `birthday`, `age`, `heigth`, `cup`, `bust`, `waist`, `hip`, `hometown`, `hobby`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	for te := range chThEt {
		fmt.Println("chThEt->", len(chThEt),"：", te.name, te.ename, te.birthday, te.age, te.heigth, te.cup, te.bust, te.waist, te.hip, te.hometown, te.hobby)
		_, e := DB.Exec(sql, te.name, te.ename, te.birthday, te.age, te.heigth, te.cup, te.bust, te.waist, te.hip, te.hometown, te.hobby)
		if e != nil {
			utils.HandleError(e, "DB.Exec ARTIST")
			fmt.Println("XXXX：插入失败：", te.name)
		}
	}
}

//错误信息存储
func insertFail()  {
	sql := "INSERT INTO `fail` (THID,STAGE,REASON) VALUES (?, ?, ?);"
	for val := range chFail {
		_, e := DB.Exec(sql, val.fid, val.stage, val.reason)
		if e != nil {
			utils.HandleError(e, "DB.Exec FAIL")
			fmt.Println("XXXXX：插入失败：", val)
		}
	}
}
