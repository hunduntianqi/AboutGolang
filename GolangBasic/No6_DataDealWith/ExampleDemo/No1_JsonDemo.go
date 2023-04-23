package main

import (
	"encoding/json"
	"fmt"
)

// JsonTest 定义Json测试结构体
type JsonTest struct {
	Company string `json:"公司"`
	Addr    string `json:"地址"`
	Persons int    `json:",string"` // 将原来为整型的字段Persons转换为字符串输出
}

func main() {
	/*
		通过结构体生成json
	*/
	// 结构体变量初始化
	var js1 JsonTest = JsonTest{Company: "富士康", Addr: "河南省郑州市航空港区综合保税区", Persons: 100000}
	// 调用json.Marshal函数生成json格式数据
	jsResult1, error1 := json.Marshal(js1)
	if error1 == nil {
		fmt.Println(string(jsResult1))
	}

	// 调用json.MarshalIndent函数格式化输出json
	jsResult2, error2 := json.MarshalIndent(js1, ``, ` `)
	if error2 == nil {
		fmt.Println(string(jsResult2))
	}
	// 解析json到结构体
	// 定义结构体变量
	var jsDecode JsonTest
	errDecode := json.Unmarshal(jsResult1, &jsDecode)
	if errDecode == nil {
		fmt.Println(jsDecode)
	} else {
		fmt.Println(errDecode)
	}

	/*
		通过map生成json
	*/
	// 定义一个map
	jsonMap := make(map[string]interface{})
	jsonMap["姓名"] = `郭鹏涛`
	jsonMap["性别"] = `男`
	jsonMap["公司"] = "富士康"
	jsonMap["地址"] = "河南省郑州市航空港区综合保税区"
	jsonMap["人数"] = "10000~300000"
	jsonMap["所属部门"] = "idpbg事业群-MLB事业处-测试工程处-失效分析工程部-维修改善三课-RF组"
	// 调用json.Marshal()函数生成json数据
	jsResult3, err3 := json.Marshal(jsonMap)
	if err3 == nil {
		fmt.Println(string(jsResult3))
	}
	// 格式化输出
	jsResult4, err4 := json.MarshalIndent(jsonMap, ``, ` `)
	if err4 == nil {
		fmt.Println(string(jsResult4))
	}
	// 解析json到map
	// 定义map数据变量
	var mapDecode map[string]interface{}
	errMapDecode := json.Unmarshal(jsResult3, &mapDecode)
	if errMapDecode == nil {
		fmt.Println("人数=", mapDecode["人数"])
	}
}
