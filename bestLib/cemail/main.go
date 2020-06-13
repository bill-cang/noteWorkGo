package main

import (
	"github.com/jordan-wright/email"
	"log"
	"mime"
	"net/smtp"
	"net/textproto"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan *email.Email, 4)
	err := sendEmailByPool(ch, &wg)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		ch <- &email.Email{
			To:      []string{"polaris@studygolang.com"},
			From:    mime.QEncoding.Encode("UTF-8", "Go语言中文网") + "<274768166@qq.com>",
			Subject: "Pool" + strconv.Itoa(i),
			HTML:    []byte("<h1>这是 HTML 正文</h1>"),
			Headers: textproto.MIMEHeader{},
		}
	}
	wg.Wait()
	close(ch)
}

func sendEmailByPool(ch <-chan *email.Email, wg *sync.WaitGroup) error {
	p, err := email.NewPool(
		"smtp.qq.com:25",
		4,
		smtp.PlainAuth("", "274768166@qq.com", "password", "smtp.qq.com"))
	if err != nil {
		return err
	}
	for i := 0; i < 4; i++ {
		go func() {
			for e := range ch {
				err := p.Send(e, 10e9)
				if err != nil {
					log.Println("Send Email fail, err:", err)
				} else {
					log.Println("Send Email Successfully!")
				}
				wg.Done()
			}
		}()
	}
	return nil
}
