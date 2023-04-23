package No6_DataDealWith

/*
	strings 包:
		实现了用于操作字符串的简单函数
		其他(不知道该分到哪一类):
			1. func Count(s, sep string) int: 返回字符串 's' 包含字符串 'sep' 的个数
			2. func Repeat(s string, count int) string: 将字符串 's' 复制 count 次后返回
		判断相关:
			1. func EqualFold(s, t string) bool: 判断两个字符串内容是否相同, 忽略大小写
			2. func HasPrefix(s, prefix string) bool: 判断字符串 's' 是否以字符串 'prefix' 开头
			3. func HasSuffix(s, suffix string) bool: 判断字符串 's' 是否以字符串 'suffix' 结尾
			4. func Contains(s, substr string) bool: 判断字符串 's' 是否包含 字符串 'substr'
			5. func ContainsRune(s string, r rune) bool: 判断字符串 's' 是否包含 字符 'r'
			6. func ContainsAny(s, chars string) bool: 判断字符串 's' 是否包含字符串 'chars' 中的任一字符
		索引相关:
			1. func Index(s, sep string) int: 返回字符串 's' 中字符串 'sep' 第一次出现位置的索引值, 没有则返回-1
			2. func IndexByte(s string, c byte) int: 返回字符串 's' 中字符 'c' 第一次出现位置的索引值, 没有则返回-1
			3. func IndexRune(s string, r rune) int: 返回字符串 's' 中字符 'r' 第一次出现位置的索引值, 没有则返回-1
			4. func IndexAny(s, chars string) int: 返回字符串 's' 中字符串 'chars' 任一字符第一次出现位置的索引值, 没有或 'chars' 为空则返回-1
			5. func LastIndex(s, sep string) int: 返回字符串 's' 中字符串 'sep' 最后出现位置的索引值, 没有则返回-1
			6. func LastIndexAny(s, chars string) int: 返回字符串 's' 中字符串 'chars' 任一字符最后出现位置的索引值, 没有或 'chars' 为空则返回-1
		大小写转换相关:
			1. func Title(s string) string: 将字符串 's' 中每个单词首字母转换为大写后返回
			2. func ToLower(s string) string: 将字符串 's' 中所有字母转换为小写后返回
			3. func ToLowerSpecial(_case unicode.SpecialCase, s string) string: 根据 '_case' 规定的字符映射, 将字符串 's' 中所有字母转换为小写后返回
			4. func ToUpper(s string) string: 将字符串 's' 中所有字母转换为大写后返回
			5. func ToUpperSpecial(_case unicode.SpecialCase, s string) string: 根据 '_case' 规定的字符映射, 将字符串 's' 中所有字母转换为大写后返回
			6. func ToTitle(s string) string: 将字符串 's' 中所有字母转换为大写后返回
			7. func ToTitleSpecial(_case unicode.SpecialCase, s string) string: 根据 '_case' 规定的字符映射, 将字符串 's' 中所有字母转换为大写后返回
		替换和去除字符相关:
			1. func Replace(s, old, new string, n int) string: 将字符串 's' 中的字符串 'old' 替换为字符串 'new', 'n' 为替换次数, 若 'n < 0', 则替换所有
																符合条件的字符串
			2. func Trim(s string, cutset string) string: 去除字符串 's' 前后的字符串 'cutest' 后返回
			3. func TrimSpace(s string) string: 去除字符串 's' 前后的空白字符后返回, 包括换行符, 制表符等
			4. func TrimFunc(s string, f func(rune) bool) string: 去除字符串 's' 前后的符合函数 'f' 的字符后返回
			5. func TrimLeft(s string, cutset string) string: 去除字符串 's' 前端的字符串 'cutest' 后返回
			6. func TrimLeftFunc(s string, f func(rune) bool) string: 去除字符串 's' 前端的符合函数 'f' 的字符后返回
			7. func TrimPrefix(s, prefix string) string: 去除字符串 's' 前端的字符串 'prefix' 后返回
			8. func TrimRight(s string, cutset string) string: 去除字符串 's' 后端的字符串 'cutest' 后返回
			9. func TrimRightFunc(s string, f func(rune) bool) string: 去除字符串 's' 后端的符合函数 'f' 的字符后返回
			10. func TrimSuffix(s, suffix string) string: 去除字符串 's' 后端的字符串 'suffix' 后返回
		字符串分割:
			1. func Fields(s string) []string: 将字符串 's' 按照空白字符分割后返回一个字符串切片
			2. func FieldsFunc(s string, f func(rune) bool) []string: 将字符串 's' 按照符合函数 'f' 的字符分割后返回一个字符串切片
			3. func Split(s, sep string) []string: 将字符串按照指定的分隔符分割后返回字符串切片, 如果 'sep' 为空字符, 会将s切分成每一个unicode码值一个字符串
			4. func SplitN(s, sep string, n int) []string: 将字符串按照指定的分隔符分割后返回字符串切片, 参数 'n' 决定返回切片中数据个数
					n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
					n == 0: 返回nil
					n < 0 : 返回所有的子字符串组成的切片
			5. func SplitAfter(s, sep string) []string: 将字符串按照指定的分隔符, 从分隔符后面分割后返回字符串切片
			6. func SplitAfterN(s, sep string, n int) []string: 将字符串按照指定的分隔符, 从分隔符后面分割后返回字符串切片, 参数 'n' 决定返回切片中数据个数
					n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
					n == 0: 返回nil
					n < 0 : 返回所有的子字符串组成的切
		字符串拼接:
			func Join(a []string, sep string) string: 将字符串切片 'a' 中的字符串按照分隔符 'sep' 连接起来后返回字符串
*/
