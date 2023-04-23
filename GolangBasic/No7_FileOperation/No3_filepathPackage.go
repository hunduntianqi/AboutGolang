package No7_FileOperation

/*
	filepath 包:
		filepath包实现了兼容各操作系统的文件路径相关的实用操作函数
		常用函数:
			1. func IsAbs(path string) bool: 判断指定路径是否是一个绝对路径
			2. func Abs(path string) (string, error): 返回指定路径地址代表的绝对路径, 如果path不是绝对路径, 会加入当前工作目录
														使之成为绝对路径
			3. func Join(elem ...string) string: 将任意数量的路径元素以路径分隔符拼接为一个单一路径返回
			4. func FromSlash(path string) string: 将指定路径中的斜杠（'/'）替换为路径分隔符并返回替换结果
			5. func ToSlash(path string) string: 将指定路径的路径分隔符替换为斜杠（'/'）并返回替换结果
			6. func Glob(pattern string) (matches []string, err error): 返回所有符合匹配字符串pattern的文件名称
				例: 获取某个文件夹下所有文件的绝对路径 ==> filepath.Glob(dirPath + "//*")
*/
