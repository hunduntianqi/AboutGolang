package main

/*
   fmt包:
      主要分为向外输出内容和获取输入内容两部分
       向外输出:
           Print系列函数:
               Print函数: 直接输出内容
               Printf函数: 支持格式化输出内容
               Println函数: 会在输出内容末尾添加一个换行符
           Fprint系列函数:
               将内容输出到一个io.Writer接口类型的变量w中, 我们通常用这个函数往文件中写入内容
               func Fprint(w io.Writer, a ...interface{}) (n int, err error)
               func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
               func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
               注意: 只要满足io.Writer接口的类型都支持写入
           Sprint系列函数:
               会把传入的数据生成并返回一个字符串
               func Sprint(a ...interface{}) string
               func Sprintf(format string, a ...interface{}) string
               func Sprintln(a ...interface{}) string
           Errorf函数:
               根据format参数生成格式化字符串并返回一个包含该字符串的错误
               func Errorf(format string, a ...interface{}) error
               该函数通常被用来自定义错误类型
       获取输入:
           fmt.Scan函数:
               func Scan(a ...interface{}) (n int, err error)
           fmt.Scanf函数:
               func Scanf(format string, a ...interface{}) (n int, err error)
           fmt.Scanln函数:
               func Scanln(a ...interface{}) (n int, err error)
           Fscan系列函数:
               功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数, 只不过它们不是从标准输入中读取数据
               而是从io.Reader中读取数据
               func Fscan(r io.Reader, a ...interface{}) (n int, err error)
               func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
               func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
           Sscan系列:
               功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数, 只不过它们不是从标准输入中读取数据
               而是从指定字符串中读取数据
               func Sscan(str string, a ...interface{}) (n int, err error)
               func Sscanln(str string, a ...interface{}) (n int, err error)
               func Sscanf(str string, format string, a ...interface{}) (n int, err error)
*/

func main() {

}
