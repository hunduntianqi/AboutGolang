package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
	使用Soup第三方库解析页面数据, 类似于 Python 的 bs4
		安装第三方库 ==> go get github.com/anaskhan96/soup
		创建解析对象 ==> soupHtml := soup.HTMLParse(html)
		获取单个标签对象 ==> tag := soupHtml.Find("标签名", "属性名", "属性值") ==> 获取一个满足条件的标签对象
		获取多个标签对象 ==> tagList := soupHtml.FindAll("标签名", "属性名", "属性值") ==> 获取一个满足条件的标签对象的列表
			注意 ==> 标签对象还可以通过Find() / FindAll()方法获取标签下阶满足条件的标签对象
		获取标签文本内容 ==> tag.Text() ==> 只能获取自己的文本内容 / tag.FullText() ==> 可获取标签及其嵌套标签所有的文本数据
		获取标签属性值 ==> tag.Attrs()["属性名"]
*/

func main() {
	// 创建文件夹存储数据
	os.MkdirAll("No16_Spider/4K", 0777)
	// 创建客户端对象
	client := &http.Client{}
	// 创建map对象, 存储请求头数据
	headers := make(map[string]string)
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	// 定义计数器, url
	num := 1
	url := ""
	// 循环下载数据
	for {
		if num == 1 {
			url = "https://pic.netbian.com/4kmeinv/"
		} else {
			url = "https://pic.netbian.com/4kmeinv/index_" + strconv.Itoa(num) + ".html"
		}
		// 创建请求对象
		request, _ := http.NewRequest("GET", url, nil)
		// UA伪装
		request.Header.Set("user-agent", headers["user-agent"])
		// 发起请求获取响应数据
		response, _ := client.Do(request)
		// 转换响应数据
		data, _ := io.ReadAll(response.Body)
		html := string(data)
		// 转换编码为utf-8
		html = encodingString(html)
		//fmt.Println(html)
		// 使用soup第三方库获取解析对象
		soupHtml := soup.HTMLParse(html)
		// 解析数据
		divSlist := soupHtml.Find("div", "class", "slist").FindAll("li")
		//fmt.Println(divSlist)
		for _, li := range divSlist {
			picHtmlUrl := "https://pic.netbian.com" + li.Find("a").Attrs()["href"]
			picName := li.Find("b").Text()
			// 调用函数, 解析获取大图链接
			picUrl := getBigPicUrl(picHtmlUrl, headers, client)
			// 调用函数, 下载图片
			downPic(picUrl, picName, client, headers)
			fmt.Printf("%s ==> %s下载完毕!!\n", picName, picUrl)
		}
		fmt.Printf("第%d页数据抓取完毕\n", num)
		num++
		response.Body.Close()
		if strings.Contains(soupHtml.Find("div", "class", "page").FullText(), "下一页") {
			fmt.Println("存在下一页数据, 继续抓取!!")
		} else {
			log.Fatal("数据抓取完毕, 退出程序")
		}
	}
}

// 定义函数, 转换文本编码格式为utf-8
func encodingString(htmlString string) string {
	htmlString, _ = simplifiedchinese.GBK.NewDecoder().String(htmlString)
	return htmlString
}

// 定义函数, 解析获取大图Url
func getBigPicUrl(url string, headers map[string]string, client *http.Client) string {
	// 创建请求对象
	req, _ := http.NewRequest("GET", url, nil)
	// 构建请求头
	req.Header.Set("user-agent", headers["user-agent"])
	// 发送请求
	response, _ := client.Do(req)
	htmlData, _ := io.ReadAll(response.Body)
	// 创建soup对象
	soupHtml := soup.HTMLParse(string(htmlData))
	// 解析获取大图链接
	bigPicUrl := "https://pic.netbian.com" + soupHtml.Find("a", "id", "img").Find("img").Attrs()["src"]
	defer response.Body.Close()
	return bigPicUrl
}

// 定义函数, 下载图片并保存
func downPic(url, picName string, client *http.Client, headers map[string]string) {
	// 创建请求对象
	reqPic, _ := http.NewRequest("GET", url, nil)
	// 构造UA伪装
	reqPic.Header.Set("user-agent", headers["user-agent"])
	// 发送请求
	response, _ := client.Do(reqPic)
	picData, _ := io.ReadAll(response.Body)
	// 创建文件对象
	file, _ := os.Create("No16_Spider/4K/" + picName + ".jpg")
	// 写入数据
	file.Write(picData)
	// 关闭文件
	file.Close()
	defer response.Body.Close()
}
