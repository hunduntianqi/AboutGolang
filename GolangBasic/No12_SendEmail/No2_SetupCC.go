package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

/*
	设置抄送功能:
		设置抄送邮箱 ==> emailName.Cc = []String{"emailAddress1", "emailAddress2", ...}
		设置秘密抄送邮箱 ==> emailName.Bcc = []String{"emailAddress1", "emailAddress2", ...}
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
	emailName.Subject = "Golang发送邮件测试程序"
	// 设置邮件正文内容
	emailName.Text = []byte("Golang邮件测试程序发送成功提示!!!")
	// 发送邮件
	errSend := emailName.Send("smtp.qq.com:25", smtp.PlainAuth("", "1729992141@qq.com",
		"kbxwniucpokcchgf", "smtp.qq.com"))
	if errSend != nil {
		fmt.Println("邮件发送失败 ==> ", errSend)
		return
	}
	fmt.Println("邮件发送成功!!")
}
