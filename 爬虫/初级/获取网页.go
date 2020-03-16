package main

/*
爬取个人作品信息
*/
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"utils"
)

var (
	filePor *os.File
	fileAdd string
	chqu    = make(chan string, 1)
)
//★★★核心：从老师管道读取各位老师地址，并按老师id发起二次请求,获取老师个人信息放入老师管道，获取老师作品放入作品管道
func getPageDataToChan() {
	var url string
	var errThid string //发生错的老师的id
	//var starPag int = 1 //预定义开始页，发生错误时，修改该值，实现端点续爬
	//异常处理，断电续续爬
	defer func() {
		if err := recover(); err != nil {
			errStr := fmt.Sprintf("%s", err)
			if strings.Index(errStr, err001) > 0 {
				//DB.Exec(sql_fail, errStr, "请求个人主页", err)
				fail := &fails{errThid, "请求个人主页", errStr}
				chFail <- fail
				chThid <- errThid
				fmt.Println("====================", errThid, "触发休眠！！！")
				time.Sleep(20 * time.Second)
				go getPageDataToChan()
			}
		}
	}()

	//读取管道里的老师编号
	for lsStr := range chThid {
		go func(lsStr string) {
			errThid = lsStr
			lsNum, e := strconv.Atoi(lsStr)
			var thNmae string
			if e != nil {
				utils.HandleError(e, "strconv.Atoi")
			}
			url = fmt.Sprintf("https://www.9fanhao.com/special-read-id-%d-p-1.html", lsNum)
			pageStr := getPageStr(url)

			compile0 := regexp.MustCompile(repxls_name)
			ls_nameSpl := compile0.FindAllStringSubmatch(pageStr, 1)
			compile1 := regexp.MustCompile(repxls_mass)
			ls_massSpl := compile1.FindAllStringSubmatch(pageStr, 1)
			compile2 := regexp.MustCompile(repxls_copa)
			ls_conpaSpl := compile2.FindAllStringSubmatch(pageStr, 1)
			//fmt.Printf("=============%v老师共有作品%v页\n", ls_nameSpl[0][1], ls_conpaSpl[0][1])
			thNmae = ls_nameSpl[0][1]

			//开辟一条协程去单独整理老师个人信息
			go getArtistMsg(ls_nameSpl, ls_massSpl)

			ln:= len(ls_conpaSpl)
			if ln < 1{
				return
			}

			maxPage := aToI(ls_conpaSpl[0][1]) + 1
			for i := 1; i < maxPage; i++ {
				url := fmt.Sprintf("https://www.9fanhao.com/special-read-id-%d-p-%d.html", lsNum, i)
				response, err := http.Get(url)
				utils.HandleError(err, "http.Get")
				defer response.Body.Close()
				bytes, err := ioutil.ReadAll(response.Body)
				utils.HandleError(err, "ioutil.ReadAll")
				pageStr = string(bytes)

				//番号信息
				rexPor := regexp.MustCompile(repxFh)
				submatch := rexPor.FindAllStringSubmatch(pageStr, -1)

				for _, val := range submatch {
					//fmt.Println(tname,"老师第",i,"页作品：", val[1], "\t", val[2], "\t", val[3], "\t", val[4])
					actressPor := new(video)
					actressPor.code, actressPor.title, actressPor.issueData, actressPor.filmlength = val[1], val[2], val[3], val[4]
					actressPor.teacher = thNmae
					chVido <- actressPor
				}
				time.Sleep(5 * time.Nanosecond)
			}
		}(lsStr)
	}
	time.Sleep(30*time.Second)
	chExit <- StarPage
}

//解析老师们的个人信息放入管道
func getArtistMsg(nameSpl [][]string, masgSpl [][]string) {
	bg, ed := 0, 0
	thp := new(teacher)
	thp.name = nameSpl[0][1]
	masgStr := masgSpl[0][0]

	if strings.Index(masgStr, "英文名字") > -1 {
		bg = strings.Index(masgStr, "：") + 3
		ed = strings.Index(masgStr, "</p>")
		thp.ename = masgStr[bg:ed]
	}
	//去掉所有空格
	masgStr = utils.CompressStr(masgStr)

	if strings.Index(masgStr, "生日:") > -1 {
		bg = strings.Index(masgStr, "生日:") + len("生日:")
		ed = bg + 10
		thp.birthday = masgStr[bg:ed]
	} else {
		thp.birthday = "9999-07-21"
	}
	if strings.Index(masgStr, "年龄:") > -1 {
		bg = strings.Index(masgStr, "年龄:") + len("年龄:")
		ed = bg + 2
		thp.age = aToI(masgStr[bg:ed])
	}
	if strings.Index(masgStr, "身高:") > -1 {
		bg = strings.Index(masgStr, "身高:") + len("身高:")
		ed = bg + 3
		thp.heigth = aToI(masgStr[bg:ed])
	}
	if strings.Index(masgStr, "罩杯:") > -1 {
		bg = strings.Index(masgStr, "罩杯:") + len("罩杯:")
		ed = bg + 1
		thp.cup = masgStr[bg:ed]
	}
	if strings.Index(masgStr, "胸围:") > -1 {
		bg = strings.Index(masgStr, "胸围:") + len("胸围:")
		ed = bg + 2
		thp.bust = aToI(masgStr[bg:ed])
	}
	if strings.Index(masgStr, "腰围:") > -1 {
		bg = strings.Index(masgStr, "腰围:") + len("腰围:")
		ed = bg + 2
		thp.waist = aToI(masgStr[bg:ed])
	}
	if strings.Index(masgStr, "臀围:") > -1 {
		bg = strings.Index(masgStr, "臀围:") + len("臀围:")
		ed = bg + 2
		thp.hip = aToI(masgStr[bg:ed])
	}
	if strings.Index(masgStr, "出生地:") > -1 {
		bg = strings.Index(masgStr, "出生地:") + len("出生地:")
		masgStr = masgStr[bg:]
		ed = strings.Index(masgStr, "</p>")
		thp.hometown = masgStr[:ed]
	}
	if strings.Index(masgStr, "爱好:") > -1 {
		bg = strings.Index(masgStr, "爱好:") + len("爱好:")
		ed = strings.LastIndex(masgStr, "</p>")
		thp.hobby = masgStr[bg:ed]
	}
	//fmt.Println(thp.name, ",", thp.ename, ",", thp.birthday, ",", thp.age, ",", thp.heigth, ",", thp.cup, ",", thp.bust, ",", thp.waist, ",", thp.hip, ",", thp.hometown, ",", thp.hobby)
	//将老师信息固化到数据库
	//insertArtist(thp)
	chThEt <- thp
	//fmt.Printf("%v老师个人信息爬取完毕已存放信息管道！(%d)\n", thp.name, len(chThEt))
}

//字符串转数字处理
func aToI(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		fmt.Println(str)
		utils.HandleError(e,"strconv.Atoi(str)")
		return 0
	}
	return i
}

//写出数据到本地文件主函数
/*func outPutFilemain() {
	var fileName string
	go func() {
		for i := 1; i < 17; i++ {
			url := fmt.Sprintf("https://www.9fanhao.com/special-read-id-14123-p-%d.html", i)
			pageStr := getPageStr(url)

			//老师名字
			if i == 1 {
				lsName := regexp.MustCompile(repxXm).FindAllStringSubmatch(pageStr, -1)
				fileName = lsName[0][1]
			}

			//番号信息
			rexPor := regexp.MustCompile(repxFh)
			submatch := rexPor.FindAllStringSubmatch(pageStr, -1)
			//fmt.Printf("共找到：%v个\n", len(submatch))
			count := 1
			for _, val := range submatch {
				fmt.Println(count, "：", val[1], "\t", val[2], "\t", val[3], "\t", val[4])
				str := val[1] + "\t" + val[2] + "\t" + val[3] + "\t" + val[4] + "\n"
				chdata <- str
				count++
			}

		}
	}()

	fileAdd = "E:/workspace/Goland/爬虫/初级/file/" + fileName + ".txt"
	go writeToFIle()
	<-chqu
}*/

//获取网页数据
func getPageStr(url string) (pageStr string) {
	response, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	utils.HandleError(err, "ioutil.ReadAll")
	pageStr = string(bytes)
	return
}

//写出文件
/*func writeToFIle() {
	filePro, e := os.OpenFile(fileAdd, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0754)
	utils.HandleError(e, "os.Open")
	writer := bufio.NewWriter(filePro)

	defer filePor.Close()
	count := 1
	for cone := range chdata {
		fmt.Println("*", count, "：", cone)
		writer.WriteString(strconv.Itoa(count) + ":" + cone)
		count++
	}
	writer.Flush()
	chqu <- "over"

}*/
