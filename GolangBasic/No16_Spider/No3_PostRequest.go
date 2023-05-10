package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

/*
	发送POST请求:
		1. 创建客户端对象 ==> client :=&http.Client{}
		2. 封装参数:
			a. 创建 map 或 结构体 对象, 将要提交的表单参数以键值对形式存入map 或 结构体
			b. 将 map 或 结构体 对象序列化为 json 字符串 ==> postParamStr, err := json.Marshal( map / struct )
			c. 根据json字符串创建 io.Reader 对象 ==> reader := strings.NewReader(postParamStr)
		3. 创建请求对象 ==> request, err := http.NewRequest("POST", url, reader) ==> 注意: 请求的方法 POST 必须大写
		4. 设置请求头参数:
			对于POST请求必须设置文本类型 ==>  request.Header.Set("Content-Type", "实际的数据类型")
				数据类型例子 ==> "application/json", "text/plain", "*\*"(备注: 此处是斜杠/, 因注释问题加了个\斜杠)
			添加 ua 伪装 ==> request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
			添加Cookie ==> request.Header.Add("Cookie", "SID=VghNiAxPzak-sLGoiCQJx5jIa9XJXF1saDebjBUgLROL7bhjyjwGszQDW9dZFJ_556t3Cg.; __Secure-1PSID=VghNiAxPzak-sLGoiCQJx5jIa9XJXF1saDebjBUgLROL7bhjiFvrfclEyryvfo-pV62dYw.; __Secure-3PSID=VghNiAxPzak-sLGoiCQJx5jIa9XJXF1saDebjBUgLROL7bhjOiEno0puzW5ZUVmPzU0a6Q.; HSID=A5XlDHEvA3kwCsWTh; SSID=AmGVvHVfdZftdUgdY; APISID=d81ntH2M4QoZxWfe/A8p5ubBlTDkKfdzLn; SAPISID=miETsMny_7bPMCGq/A7ONMZ8zeEolmZ5N2; __Secure-1PAPISID=miETsMny_7bPMCGq/A7ONMZ8zeEolmZ5N2; __Secure-3PAPISID=miETsMny_7bPMCGq/A7ONMZ8zeEolmZ5N2; _ga=GA1.3.1522522022.1681909947; OTZ=6993432_24_24__24_; SEARCH_SAMESITE=CgQIipgB; NID=511=ZGQMp9NVNC00AW1BYd-kwm5yCycVk8gAjLUSOlNht60VRSJEsQAY17pQhSV6Rq8mA2JuD0PltHirlhniIF2QPcJ-iJ5nm18LQRQQi2h1U3Co7CX1ywlb0wdw8aO3Qg8oNWb-ScL4EBdZC4s9ozNttwMyrm3ue-fQQ8moYAj4ssQamBzxu1QAMyevy-sM7qECtBWHL5U29n7niJcs180gAPHcoMzeZg5G5JafMUMJ2WkKXxnmEptrGOHr_TLM4jDPGTB-WlZvpo4C0fDQY1ubA-M0sA; AEC=AUEFqZeZ_Nzm2J2Hyh6SafatRdPhWlNIDqyCct9O2VK3EDFEwMrwr9Fsr6A; 1P_JAR=2023-05-08-12; SIDCC=AP8dLtxeV2oXwtswImvelOZr7030QlEZamzV2jUYxJvxydXeEm2S_9eHr2dfN0Hc077CXflH9A; __Secure-1PSIDCC=AP8dLty6CE54Ob398g-7qx84iRzetWWyyzYbNXtgj7GBh3D4zj8FANmnHJVaMxTIEvJV0UmcEA; __Secure-3PSIDCC=AP8dLtw5emB10FeSh2JWhKlznMglnQK2x4TdzAkWrX8RBdr4k4x0OU2ObcR6WtjnBcUL6rh0O5s")
		5. 发出请求, 获取响应对象 ==> response, err := client.Do(request)
		6. 获取响应状态码 ==> response.StatusCode
		7. 获取响应主体数据 ==> data, err := io.ReadAll(response.Body)
		8. 关闭响应主体 ==> response.Body.Close()
*/

func main() {
	// 创建客户端对象
	client := &http.Client{}
	// 创建Map, 存储表单参数
	mapParam := make(map[string]string)
	// 存入数据到map中
	mapParam["appkey"] = "0WEB0OEX9Y4SQ244"
	// map序列化为字符串并创建*Reader对象
	paramStr, _ := json.Marshal(mapParam)
	reader := strings.NewReader(string(paramStr))
	// 创建POST请求对象
	request, err := http.NewRequest("POST", "https://otheve.beacon.qq.com/analytics/v2_upload?appkey=0WEB0OEX9Y4SQ244", reader)
	if err != nil {
		log.Fatal("请求对象创建失败 ==> ", err)
	}
	// 构建请求头
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	request.Header.Set("Content-Type", "text/plain")
	// 发送请求
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("响应数据获取失败 ==> ", err)
	}
	fmt.Println("StatusCode ==> ", response.StatusCode)
	// 读取响应数据
	data, _ := io.ReadAll(response.Body)
	fmt.Println(string(data))
	// 解析json数据
	// 创建map
	parseData := make(map[string]interface{})
	// 解析数据到map中
	json.Unmarshal(data, &parseData)
	fmt.Println(parseData)
	for key, value := range parseData {
		if key == "result" {
			fmt.Println("key ==>", key, "; value ==> ", value.(float64))
		} else {
			fmt.Println("key ==>", key, "; value ==> ", value)
		}

	}
	defer response.Body.Close()
}
