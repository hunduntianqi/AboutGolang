package No6_DataDealWith

/*
	strconv包:
		实现了基本数据类型与其字符串表示的相互转换
		Parse系列 ==> 字符串转换为基本数据类型:
			1. func ParseBool(str string) (value bool, err error): 返回字符串表示的bool值, 接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
			2. func ParseInt(s string, base int, bitSize int) (i int64, err error): 返回字符串表示的整数值, 接受正负号
				base ==> 指定进制, 2~36, 如果为 0, 则会从字符串前置判断("0x"是16进制, "0"是8进制, 否则为10进制)
				bitSize ==> 指定结果必须能无溢出赋值的整数类型, 0、8、16、32、64 分别代表 int、int8、int16、int32、int64
			3. func ParseUint(s string, base int, bitSize int) (n uint64, err error): 返回字符串表示的整数值, 不接受正负号
				base ==> 指定进制, 2~36, 如果为 0, 则会从字符串前置判断("0x"是16进制, "0"是8进制, 否则为10进制)
				bitSize ==> 指定结果必须能无溢出赋值的整数类型, 0、8、16、32、64 分别代表 int、int8、int16、int32、int64
			4. func ParseFloat(s string, bitSize int) (f float64, err error): 返回字符串表示的浮点数值
				bitSize ==> 32 >> float32, 64 >> float64
		Format系列 ==> 基本数据类型转换为字符串:
			1. func FormatBool(b bool) string: 根据b的值返回 "true" 或 "false"
			2. func FormatInt(i int64, base int) string: 根据指定进制, 返回整数 i 的字符串表示, 结果中会使用小写字母 'a' 到 'z' 表示大于10的数字
				base ==> 指定进制, 2~36
			3. func FormatUint(i uint64, base int) string: 根据指定进制, 返回整数 i 的字符串表示, 为无符号整型, 结果中会使用小写字母 'a' 到 'z'
															表示大于10的数字; base ==> 指定进制, 2~36
			4. func FormatFloat(f float64, fmt byte, prec, bitSize int) string: 返回浮点数f的字符串表示形式
				fmt ==> 表示格式
					'f' ==> -ddd.dddd
					'b' ==> -ddddp±ddd, 指数为二进制
					'e' ==> -d.dddde±dd, 十进制指数
					'E' ==> -d.ddddE±dd, 十进制指数
					'g' ==> 指数很大时用 'e' 格式, 否则 'f' 格式
					'G' ==> 指数很大时用 'E' 格式, 否则 'f' 格式
				prec ==> 控制精度, 对'f'、'e'、'E', 表示小数点后数字个数; 对'g'、'G', 表示总的数字个数; 如果prec 为-1, 则代表使用最少数量的, 但又必需的数字来表示f
		Atoi & Itoa:
			Atoi ==> func Atoi(s string) (i int, err error), 等价于 ParseInt(s, 10, 0)
			Itoa ==> func Itoa(i int) string, 等价于 FormatInt(i, 10)
*/
