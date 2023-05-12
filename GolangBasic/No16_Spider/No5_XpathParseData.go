package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
	Go语言Xpath解析数据:
		安装第三方库 ==> go get github.com/antchfx/htmlquery
		创建解析对象 ==> document, err := htmlquery.Parse(resp.Body)
		获取所有对应标签对象列表 ==> nodes := htmlquery.Find(document, "xpath")
		获取单个对应标签 ==> node := htmlquery.FindOne(document, "xpath")
		获取对应标签对象属性 ==> tagAttr := htmlquery.FindOne(node, "xpath/@AttrName")
			获取属性值字符串 ==> value = htmlquery.InnerText(tagAttr)
		获取对象标签文本内容 ==> text := htmlquery.FindOne(node, "xpath/text()")
			获取文本内容字符串 ==> value = htmlquery.InnerText(text)
*/

func main() {
	start := time.Now()
	// 创建客户端对象
	client := &http.Client{}
	// 定义map对象存储请求头参数
	header := make(map[string]string)
	header["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	header["cookie"] = "Hm_lvt_ce3020881c73bb20f0830ef4ed0a61fb=1683893634; __gads=ID=8c615656b6e7d1f9-2274d572b1e000c9:T=1683893651:RT=1683893651:S=ALNI_MYaLQpe0G00WeG6g1fHSV7MrnI_7g; __gpi=UID=00000bf2e61187b7:T=1683893651:RT=1683893651:S=ALNI_Mbz_vt6Z3BLXecVZ2_zG-CIXE8I_Q; Hm_lpvt_ce3020881c73bb20f0830ef4ed0a61fb=1683894087"
	header["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	// 定义变量, 记录抓取页数
	page := 1
	// 定义变量, 记录url
	url := ""
	for {
		var request *http.Request
		if page == 1 {
			url = "https://m.bcoderss.com/tag/%E7%BE%8E%E5%A5%B3/"
			// 创建请求对象
			request, _ = http.NewRequest("GET", url, nil)
		} else {
			url = "https://m.bcoderss.com/tag/%E7%BE%8E%E5%A5%B3/page/" + strconv.Itoa(page) + "/"
			////// 创建map对象, 存储表单数据
			//from := make(map[string]string)
			//// map序列化为字符串并创建 *Reader 对象
			//paramStr, _ := json.Marshal(from)
			//reader := strings.NewReader(string(paramStr))
			// 创建请求对象
			request, _ = http.NewRequest("GET", url, nil)
		}
		// 构建请求头
		request.Header.Set("user-agent", header["user-agent"])
		request.Header.Set("cookie", header["cookie"])
		request.Header.Set("accept", header["accept"])
		// 发送请求
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("发送请求失败 ==> ", err)
			return
		}
		defer response.Body.Close()
		// 使用第三方库获取xpath解析对象
		document, errDocument := htmlquery.Parse(response.Body)
		if errDocument != nil {
			fmt.Println("解析对象创建失败 ==> ", errDocument)
			return
		}
		// 获取图片li标签对象列表
		nodes := htmlquery.Find(document, "//*[@id=\"main\"]/li")
		fmt.Println(nodes)
		// 遍历获取每个li标签对象
		for index, li := range nodes {
			// 获取大图页面链接
			bigLink := htmlquery.FindOne(li, "./a/@href")
			// 调用函数, 获取大图图片链接
			bigPicUrl := getBigUrl(htmlquery.InnerText(bigLink), client, header)
			// 获取图片名称
			picName := htmlquery.InnerText(htmlquery.FindOne(li, "./a/@title")) + "_" + strconv.Itoa(index+1)
			fmt.Printf("picName ==> %v; picBigLink ==> %v\n", picName, bigPicUrl)
		}
		fmt.Printf("第 %d 页数据抓取完毕 !!\n", page)
		if strings.Contains(htmlquery.InnerText(htmlquery.FindOne(document, "//*[@id=\"nextPage\"]/a/text()")), "加载更多") {
			fmt.Println("存在下一页数据, 继续抓取!!")
			page++
		} else {
			end := time.Now()
			// 数据抓取完毕
			log.Fatal("数据抓取完毕, 共耗时 ==> ", end.Sub(start))
		}
	}
}

// 定义函数, 解析获取大图url
func getBigUrl(link string, client *http.Client, header map[string]string) string {
	// 创建请求对象
	reqGetBigUrl, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Println("请求对象创建失败 ==> ", err)
		return ""
	}
	// 包装请求头
	reqGetBigUrl.Header.Set("user-agent", header["user-agent"])
	reqGetBigUrl.Header.Set("cookie", header["cookie"])
	reqGetBigUrl.Header.Set("accept", header["accept"])
	// 发送请求
	response, err := client.Do(reqGetBigUrl)
	if err != nil {
		fmt.Println("获取响应对象失败 ==> ", err)
		return ""
	}
	defer response.Body.Close()
	// 创建xpath解析对象
	documnet, _ := htmlquery.Parse(response.Body)
	// 解析获取大图链接
	bigPicUrl := htmlquery.InnerText(htmlquery.FindOne(documnet, "/html/body/div[4]/img/@src"))
	return bigPicUrl
}
