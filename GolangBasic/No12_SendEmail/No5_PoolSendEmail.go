package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"sync"
	"time"
)

/*
	使用连接池发送邮件
*/

func main() {
	// 定义一个管道对象, 用于存放邮件指针对象
	chanEmail := make(chan *email.Email, 10)
	// 创建一个 WaitGroup 对象用于协程同步
	var waitGroup sync.WaitGroup
	// 创建一个连接池对象
	pool, errPool := email.NewPool(
		"smtp.qq.com:25",
		10,
		smtp.PlainAuth("", "1729992141@qq.com", "jkilyqjeihoydjdd", "smtp.qq.com"),
	)
	if errPool != nil {
		fmt.Println("连接池创建失败 ==> ", errPool)
		return
	}
	// waitGroup 计数器 +10
	waitGroup.Add(10)
	// for 循环创建匿名函数, 使用连接池发送邮件
	for i := 0; i < 10; i++ {
		go func() {
			defer waitGroup.Done()
			// 从管道中取出邮件对象
			for emailName := range chanEmail {
				// 使用连接池对象发送邮件
				errSend := pool.Send(emailName, 10*time.Second)
				if errSend != nil {
					fmt.Println("连接池发送邮件失败 ==> ", errSend)
					return
				}
			}
		}()
	}

	// for 循环创建邮件对象
	for i := 0; i < 10; i++ {
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
		// 邮件对象存入管道
		chanEmail <- emailName
	}
	fmt.Println("测试提示1")
	// 阻塞主协程
	waitGroup.Wait()
	fmt.Println("测试提示2")
	// 邮件发送完毕, 关闭管道
	close(chanEmail)
	// 关闭连接池, 如果邮件发送失败, 会造成程序阻塞
	//pool.Close()
}
