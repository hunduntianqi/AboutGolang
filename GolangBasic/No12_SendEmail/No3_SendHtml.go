package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

/*
	发送 html 格式正文数据 ==> emailName.HTML = []byte(`html 格式文本数据`)
*/

func main() {
	// 创建一个邮件对象
	emailName := email.NewEmail()
	// 设置发送邮件邮箱地址
	emailName.From = "混沌天炁<1729992141@qq.com>"
	// 设置接收邮件邮箱地址, 可设置多个接收邮箱
	emailName.To = []string{"1729992141@qq.com", "风亦飘兮<2360199712@qq.com>"}
	// 设置抄送邮箱
	emailName.Cc = []string{"2360199712@qq.com"}
	// 设置秘密抄送邮箱
	emailName.Bcc = []string{"17320101759@189.cn"}
	// 设置邮件主题
	emailName.Subject = "Golang发送 HTML 数据邮件测试程序"
	// 设置邮件正文内容
	emailName.HTML = []byte(`<h1>python email send to test!!</h1>
    <h1>Golang email send to test!!</h1>
    <h2>Golang email send to test!!</h2>
    <h3>Golang email send to test!!</h3>
    <h4>Golang email send to test!!</h4>
    <h5>Golang email send to test!!</h5>
    <h6>Golang email send to test!!</h6>
    <img src="https://img.5mtb.com/c2022/03/17/he2udww1soz.jpg">
    <br>
    <a href="https://www.baidu.com">点击跳转百度首页</a>`)
	// 发送邮件
	errSend := emailName.Send("smtp.qq.com:25", smtp.PlainAuth("", "1729992141@qq.com",
		"jkilyqjeihoydjdd", "smtp.qq.com"))
	if errSend != nil {
		fmt.Println("邮件发送失败 ==> ", errSend)
		return
	}
	fmt.Println("邮件发送成功!!")
}
