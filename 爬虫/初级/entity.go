package main

type video struct {
	code       string //番号
	title      string //作品名称
	teacher    string //老师名
	issueData  string //发布时间
	filmlength string //作品大小
}

type teacher struct {
	name     string //老师姓名
	ename    string //英文名
	birthday string //生日
	age      int    //年龄
	heigth   int    //身高
	cup      string //罩杯
	bust     int    //胸围
	waist    int    //腰围
	hip      int    //臀围
	hometown string //出生地
	hobby    string //爱好
}

type fails struct {
	fid    string  //错误id
	stage  string  //阶段
	reason string  //原因
}
