package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"net/http"
	"regexp"
	"strings"
)

/*
	使用 http 包发送get请求:
		1. 创建客户端对象 ==> client :=&http.Client{}
		2. 获取一个请求对象 ==> request, err := http.NewRequest("GET", url, nil) ==> 注意: 请求的方法 GET 必须大写
		3. 设置请求头参数:
			添加 ua 伪装 ==> request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
			添加Cookie ==> request.Header.Add("Cookie", "SID=VghNiAxPzak-sLGoiCQJx5jIa9XJXF1saDebjBUgLROL7bhjyjwGszQDW9dZFJ_556t3Cg.; __Secure-1PSID=VghNiAxPzak-sLGoiCQJx5jIa9XJXF1saDebjBUgLROL7bhjiFvrfclEyryvfo-pV62dYw.; __Secure-3PSID=VghNiAxPzak-sLGoiCQJx5jIa9XJXF1saDebjBUgLROL7bhjOiEno0puzW5ZUVmPzU0a6Q.; HSID=A5XlDHEvA3kwCsWTh; SSID=AmGVvHVfdZftdUgdY; APISID=d81ntH2M4QoZxWfe/A8p5ubBlTDkKfdzLn; SAPISID=miETsMny_7bPMCGq/A7ONMZ8zeEolmZ5N2; __Secure-1PAPISID=miETsMny_7bPMCGq/A7ONMZ8zeEolmZ5N2; __Secure-3PAPISID=miETsMny_7bPMCGq/A7ONMZ8zeEolmZ5N2; _ga=GA1.3.1522522022.1681909947; OTZ=6993432_24_24__24_; SEARCH_SAMESITE=CgQIipgB; NID=511=ZGQMp9NVNC00AW1BYd-kwm5yCycVk8gAjLUSOlNht60VRSJEsQAY17pQhSV6Rq8mA2JuD0PltHirlhniIF2QPcJ-iJ5nm18LQRQQi2h1U3Co7CX1ywlb0wdw8aO3Qg8oNWb-ScL4EBdZC4s9ozNttwMyrm3ue-fQQ8moYAj4ssQamBzxu1QAMyevy-sM7qECtBWHL5U29n7niJcs180gAPHcoMzeZg5G5JafMUMJ2WkKXxnmEptrGOHr_TLM4jDPGTB-WlZvpo4C0fDQY1ubA-M0sA; AEC=AUEFqZeZ_Nzm2J2Hyh6SafatRdPhWlNIDqyCct9O2VK3EDFEwMrwr9Fsr6A; 1P_JAR=2023-05-08-12; SIDCC=AP8dLtxeV2oXwtswImvelOZr7030QlEZamzV2jUYxJvxydXeEm2S_9eHr2dfN0Hc077CXflH9A; __Secure-1PSIDCC=AP8dLty6CE54Ob398g-7qx84iRzetWWyyzYbNXtgj7GBh3D4zj8FANmnHJVaMxTIEvJV0UmcEA; __Secure-3PSIDCC=AP8dLtw5emB10FeSh2JWhKlznMglnQK2x4TdzAkWrX8RBdr4k4x0OU2ObcR6WtjnBcUL6rh0O5s")
		4. 发出请求, 获取响应对象 ==> response, err := client.Do(request)
		5. 获取响应状态码 ==> response.StatusCode
		6. 获取响应主体数据 ==> data, err := io.ReadAll(response.Body)
		7. 关闭响应主体 ==> response.Body.Close()
*/

func main() {
	// 创建客户端对象
	client := http.Client{}
	// 创建请求对象
	request, errReq := http.NewRequest("GET", "http://pic.netbian.com/tupian/31574.html", nil)
	if errReq != nil {
		fmt.Println("请求对象创建失败 ==> ", errReq)
		return
	}
	// 设置请求头参数
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	// 发送请求, 获取响应对象
	response, errResp := client.Do(request)
	// 延迟调用关闭响应主体
	defer response.Body.Close()
	if errResp != nil {
		fmt.Println("获取响应失败 ==> ", errResp)
		return
	}
	fmt.Println("StatusCode ==> ", response.StatusCode)
	// 获取响应数据字节切片
	dataByteList, _ := io.ReadAll(response.Body)
	// 转换字节切片为 Utf-8 编码
	bytesGBKList, _ := simplifiedchinese.GBK.NewDecoder().Bytes(dataByteList)
	//fmt.Println(string(bytesGBKList))
	// 使用正则表达式解析数据
	// 将"\n"换行替换掉
	html := strings.Replace(string(bytesGBKList), "\n", "", -1)
	// 创建正则表达式对象, 指定匹配规则
	reg := regexp.MustCompile(`<li><a href="(.*?)" title=".*?" target="_blank"><span><img src=".*?\.jpg" alt=".*?" /></span><b>.*?</b></a></li>`)
	// 获取所有符合条件的数据
	dataList := reg.FindAllStringSubmatch(html, -1)
	for _, data := range dataList {
		fmt.Println(data)
	}
}
