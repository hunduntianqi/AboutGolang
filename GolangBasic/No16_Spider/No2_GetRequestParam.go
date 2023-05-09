package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

/*
	带参数Get请求:
		1. 创建客户端对象 ==> client :=&http.Client{}
		2. 获取一个请求对象 ==> request, err := http.NewRequest("GET", url, nil) ==> 注意: 请求的方法 GET 必须大写
		3. 创建一个 http.Values 对象 ==> query := request.URL.Query()
		4. 添加参数信息 ==> query.Add("ParamName", "ParamValue")
		5. 参数信息编码 ==> request.URL.RawQuery = query.Encode()
		6. 配置请求头参数
		7. 发送请求, 获取数据
		8. 解析json数据:
			a. 分析json数据结构, 创建符合结构的结构体或其他类型对象(切片, map等) ==> 例: parseData := make([]map[string]interface{}, 0)
			b. 解析数据 ==> json.Unmarshal(dataByteList []byte, &parseData)
			c. 遍历打印 parseData, 验证数据是否解析成功
*/

func main() {
	// 创建客户端对象
	client := http.Client{}
	// 创建请求对象
	request, errReq := http.NewRequest("GET", "https://movie.douban.com/j/chart/top_list?", nil)
	if errReq != nil {
		fmt.Println("请求对象创建失败 ==> ", errReq)
		return
	}
	for i := 0; i < 100; i++ {
		// 创建 Values{} 对象, 存储参数信息
		query := request.URL.Query()
		query.Set("type", "11")
		query.Set("interval_id", "100:90")
		query.Set("action", "")
		query.Set("start", strconv.Itoa(i*20))
		query.Set("limit", "20")
		// 参数信息编码
		request.URL.RawQuery = query.Encode()
		fmt.Println(request.URL.String())
		// 配置请求头, 添加UA伪装
		request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
		// 发送请求
		response, errResp := client.Do(request)
		if errResp != nil {
			fmt.Println("获取响应失败 ==> ", errResp)
			return
		}
		fmt.Println("StatusCode ==> ", response.StatusCode)
		// 获取页面数据
		dataByteList, _ := io.ReadAll(response.Body)
		fmt.Println(len(dataByteList))
		if len(dataByteList) == 2 {
			log.Fatal("数据全部抓取完毕") // log.Fatal 系列函数在打印日志后会强制结束程序运行
		}
		//html := string(dataByteList)
		//fmt.Println(html)
		// 创建切片对象
		parseData := make([]map[string]interface{}, 0)
		err := json.Unmarshal(dataByteList, &parseData)
		if err != nil {
			fmt.Println()
			return
		}
		//fmt.Println(v)
		for _, data := range parseData {
			fmt.Println(data["rating"])
			fmt.Println(data["rank"])
			fmt.Println(data["cover_url"])
			fmt.Println(data["is_playable"])
			fmt.Println(data["id"])
			fmt.Println(data["types"])
			fmt.Println(data["regions"])
			fmt.Println(data["title"])
			fmt.Println(data["url"])
			fmt.Println(data["actor_count"])
			fmt.Println(data["score"])
			fmt.Println(data["actors"])
			fmt.Println(data["is_watched"])
			fmt.Println("=============================")
		}
		// 关闭响应主体
		response.Body.Close()
	}
	// 关闭客户端
	client.CloseIdleConnections()
}
