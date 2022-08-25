// ref to https://ithelp.ithome.com.tw/articles/10159486
// run:
//   go mod init example
//   go get github.com/alexcesaro/mail/gomail
//   go run hinet_sendmail_txtplain.go

package main

import (
	"log"
	"github.com/alexcesaro/mail/gomail"
)

func main() {
	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", "john@msa.hinet.net", "John")
	msg.SetHeader("To", "sam@gmail.com")
	msg.SetHeader("Subject", "Hello!")
	msg.SetBody("text/plain", "Hello Sam")
//	msg.AddAlternative("text/html", "Hello <b>Sam</b>!")
//	if err := msg.Attach("p1.jpg"); err != nil {
//		log.Println(err)
//		return
//	}
	m := gomail.NewMailer("msr.hinet.net", "john@msa.hinet.net", "1234", 25)
	if err := m.Send(msg); err != nil {
		log.Println(err)
	}
}
