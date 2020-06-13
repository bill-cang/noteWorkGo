package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"mime"
	"net/smtp"
)

func main() {
	sendEmail("测试第三方 email 库", "ckx0709@foxmai.com.com")
}
func sendEmail(subject string, tos ...string) error {
	e := email.NewEmail()
	smtpUsername := "178709230@qq.com"
	e.From = mime.QEncoding.Encode("UTF-8", "Go语言中文网") + "<178709230@qq.com>"
	e.To = tos
	e.Subject = subject
	e.HTML = []byte("<h1>HTML 正文</h1>")
	e.AttachFile("E:/tmp/mzx.jpg")
	auth := smtp.PlainAuth("", smtpUsername, "private8539byC*", "smtp.qq.com")
	err := e.Send("smtp.qq.com:25", auth)
	if err != nil {
		//log.Println("Send Mail to", strings.Join(tos, ","), "error:", err)
		fmt.Printf("Send Fail to %+v", err)
		return err
	}
	//log.Println("Send Mail to", strings.Join(tos, ","), "Successfully")
	fmt.Printf("Send Success .")
	return nil
}
