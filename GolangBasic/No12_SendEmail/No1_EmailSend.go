package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

/*
	安装第三方库 ==> go get github.com/jordan-wright/email
	邮件发送:
		1. 创建一个邮件对象 ==> emailName := email.NewEmail()
		2. 设置发送邮件邮箱地址 ==> emailName.From = "emailAddress"
		3. 设置接收邮件邮箱地址 ==> emailName.To = []string{"emailAddress1", "emailAddress2", ...}
			设置抄送邮箱地址 ==> emailName.Cc = []String{"emailAddress1", "emailAddress2", ...}
			设置秘密抄送邮箱地址 ==> emailName.Bcc = []String{"emailAddress1", "emailAddress2", ...}
		4. 设置邮件主题 ==> emailName.Subject = "主题内容"
		5. 设置邮件正文内容:
			emailName.Text = []byte(正文数据) ==> 正文数据为普通文本
			emailName.HTML = []byte(`html 格式文本数据`) ==> 正文数据为 html 格式
			emailName.AttachFile("filePath") ==> 正文为一个附件
		6. 发送邮件 ==> emailName.Send("host:port", smtp.PlainAuth(identity, username, password, host)):
			identity ==> 为空字符串即可
			username ==> 登录账户
			password ==> 授权码
			host ==> 邮箱服务器地址
			port ==> 端口号, 一般为 25 或 465
	连接池:
		每次调用Send时都会和 SMTP 服务器建立一次连接, 如果发送邮件多且频繁时, 可能会有性能问题, 连接池可以复用网络连接, 解决此问题
		创建连接池 ==> pool, errPool := email.NewPool("host:port", poolNum, smtp.PlainAuth(identity, username, password, host))
			poolNum ==> 连接池可同时发送邮件数量
			identity ==> 为空字符串即可
			username ==> 登录账户
			password ==> 授权码
			host ==> 邮箱服务器地址
			port ==> 端口号, 一般为 25 或 465
		使用连接池:
			1. 创建一个连接池 ==> pool
			2. 创建 多个邮件对象 ==> emailName1, emailName2, ...
			3. 使用连接池发送邮件 ==> pool.Send(emailName, timeout)
				emailName ==> 要发送的邮件对象
				timeout ==> 超时时间
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
