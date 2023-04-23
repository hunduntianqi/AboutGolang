package No6_DataDealWith

/*
	JSON:
		是一种比XML更轻量级的数据交换格式, 再易于人们阅读和编写的同时, 也易于程序解析和生成,
		JSOn采用完全独立于编程语言的文本格式, 表现为键值对集合的文本描述形式, 使JSON成为较为理想的
		跨平台, 跨语言的数据交换语言
		开发者可以用JSON传输简单的字符串, 数字, 布尔值, 也可以传输一个数组, 或者一个更复杂的符合结构,
		在web开发领域中, JSON被广泛应用于Web服务端和客户端之间的数据通信
    Go语言对Json的支持:
		使用Go语言内置的encoding / json标准库, 开发者可以轻松使用Go程序生成和解析JSON格式的数据
	编码json:
		1. 通过结构体生成JSON
			注意:结构体中字段名首字母必须大写, 否则无法读取字段！！
			使用json.Marshal()函数可以对一组数据进行JSON格式编码, json.Marshal()函数声明如下:
				func Marshal(v interface{}) ([]byte, error)
			格式化输出json:
				func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
			struct_tag使用:
				结构体二次编码:
					在结构体字段后加入语句:`json:"字段名"`, 则输出字段会按照此字段名输出, 此操作属于二次编码
				指定不输出字段:
					在结构体字段后加入语句:`json:"-"`
				修改指定字段数据类型:
					在结构体字段后加入语句:`json:",数据类型"`
		2. 通过map生成JSON:
			2.1 先定义一个map数据, 并初始化
			2.2 调用json.Marshal()函数或json.MarshalIndent()函数生成json数据
	解码JSON:
		func Unmarshal (data []byte, v interface{}) error {}
		1. 解码到结构体:
			func Unmarshal(daya []byte, &结构体变量) error {}
			注意: 传入结构体变量必须为取地址, 否则无法修改实际结构体变量内容
		2. 解码到map
			func Unmarshal(daya []byte, &map数据变量) error {}
			注意: 传入map数据变量必须为取地址, 否则无法修改实际结构体变量内容
*/
