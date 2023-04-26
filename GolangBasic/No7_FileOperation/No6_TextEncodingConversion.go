package No7_FileOperation

/*
	文本字符串UTF-8和GBK编码转换:
		Go语言默认编码方式为 UTF-8, 在涉及到其他编码格式的文本处理时需要进行编码转换
		安装第三方包: go get golang.org/x/text
		UTF-8 ==> GBK(编码):
			字符串转字符串:
				newStr, err := simplifiedchinese.GBK.NewEncoder().String(oldStr)
				注意: 转换失败会返回空字符串和错误描述
			字符串转字节切片:
				newByteSlice, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(oldStr))
				注意: 转换失败后会返回 nil 和错误描述
		GBK ==> UTF-8(解码):
			字符串转字符串:
				newStr, err := simplifiedchinese.GBK.NewDecoder().String(oldStr)
				注意: 转换失败会返回空字符串和错误描述
			字符串转字节切片:
				newByteSlice, err := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(oldStr))
				注意: 转换失败后会返回 nil 和错误描述
		除GBK与UTF-8转换外, simplifiedchinese包还支持GB18030、GB2312与UTF-8的转换, 与转换GBK类似
*/
