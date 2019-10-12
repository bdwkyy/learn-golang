package main

import (
	"strings"
	//"crypto/tls"
	"flag"
	"fmt"
	"os"

	"github.com/gomail"
)

var (
	eSubject    = flag.String("subject", "Server Crash ", "email subject")
	eBody       = flag.String("content", "Server Crash ", "email content")
	eAttachPath = flag.String("attachPath", "", "email attach")
	eAddrs      = flag.String("toAddrs", "yang.yang@chinashuguo.com ", "email to addrs")
)

func main() {
	/*
	   d := gomail.NewDialer("smtp.exmail.qq.com", 465, "sw.dev@chinashuguo.com", "!ShuGuo@123#")
	   d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	   var s gomail.SendCloser
	   var err error
	   if s, err = d.Dial(); err != nil {
	       panic(err)
	   }
	   if err := gomail.Send(s, "hhhhhelo world!"); err != nil {
	       log.Print(err)
	   }
	*/
	flag.Parse()
	if 3 > len(os.Args) {
		fmt.Println("Usage: emailT --help")
		return
	}

	eEmailAddr := make([]string, 5)

	fmt.Println("参数详情:")
	fmt.Println(*eSubject)
	fmt.Println(*eBody)
	fmt.Println(*eAttachPath)
	fmt.Println(*eAddrs)
	m := gomail.NewMessage()

	m.SetHeader("From", "sw.dev@chinashuguo.com")
	eEmailAddr = strings.Split(*eAddrs, ",")
	m.SetHeader("To", eEmailAddr...)
	m.SetAddressHeader("Cc", "sw.dev@chinashuguo.com", "SG")
	if *eSubject != "" {
		m.SetHeader("Subject", *eSubject)
	}
	if *eBody != "" {
		m.SetBody("text/html", *eBody)
	}
	if *eAttachPath != "" {
		m.Attach(*eAttachPath)
	}

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "sw.dev@chinashuguo.com", "PassWD")
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	// Send emails using d.
}
