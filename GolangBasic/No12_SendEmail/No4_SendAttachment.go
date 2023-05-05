package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

/*
	邮件发送附件 ==> emailName.AttachFile("filePath")
*/

func main() {
	// 创建一个邮件对象
	emailName := email.NewEmail()
	// 设置发送邮件邮箱地址
	emailName.From = "1729992141@qq.com"
	// 设置接收邮件邮箱地址, 可设置多个接收邮箱
	emailName.To = []string{"1729992141@qq.com", "2360199712@qq.com"}
	// 设置邮件主题
	emailName.Subject = "Golang发送邮件测试程序"
	// 设置邮件正文内容
	emailName.Text = []byte("附件请查收:")
	_, errFile := emailName.AttachFile("E:\\Users\\Administrator\\Desktop\\AboutGolang-main.zip")
	if errFile != nil {
		fmt.Println("未找到对应文件 ==> ", errFile)
		return
	}
	// 发送邮件
	errSend := emailName.Send("smtp.qq.com:25", smtp.PlainAuth("", "1729992141@qq.com",
		"kbxwniucpokcchgf", "smtp.qq.com"))
	if errSend != nil {
		fmt.Println("邮件发送失败 ==> ", errSend)
		return
	}
	fmt.Println("邮件发送成功!!")
}
