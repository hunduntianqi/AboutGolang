package main

/*
	GoLang简介:
       Go是一个开源的编程语言,可以很容易的构建简单, 可靠且高效的软件
       Go语言特点:
           1. 运行效率高,开发高效,部署简单
           2. 语言层面支持并发,易于利用多核实现并发
           3. 内置runtime(作用: 性能监控, GC等)
           4. 简单易学,丰富的标准库, 强大的网络库
           5. 内置强大的工具(gofmt), 跨平台编译, 内嵌C支持
       Go语言应用:
           1. 服务器编程:如处理日志,数据打包,虚拟机处理, 文件系统等
           2. 分布式系统: 数据库代理器,中间件等
           3. 网络编程: Web应用, API应用等
           4. 云平台
       Go语言命令行工具:
            go env: 用于打印Go语言的环境信息
			go run: 编译并运行命令源码文件
			go get: 根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包, 并对它们进行编译和安装
			go build: 用于编译指定的源码文件或代码包以及它们的依赖包
			go install: 用于编译并安装指定的代码包及它们的依赖包
			go clean: 删除掉执行其它命令时产生的一些文件和目录
			go doc: 打印附于Go语言程序实体上的文档
			go test: 用于对Go语言编写的程序进行测试
			go list: 列出指定的代码包的信息
			go fix: 把指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码
			go vet: 用于检查Go语言源码中静态错误的简单工具
			go tool pprof: 交互式的访问概要文件的内容
   GoLang安装:
        Go官网下载地址：https://golang.org/dl/
        配置GOPATH:
        GOPATH指一个环境变量, 用来表明Go项目的存放路径, 项目代码存放在GOPATH的src目录下
            配置路径 ==> 系统环境变量:
                我的电脑->属性->高级系统设置->环境变量: 在系统环境变量中添加GOPATH, 内容为: GO项目路径
        GOROOT: 配置Golang编译器, 路径为Golang安装路径
        GOPATH目录下存放三个文件夹:
            bin: 用来存放编译后生成的可执行文件
            pkg: 用来存放编译后生成的归档文件
            src: 用来存放源码文件
    GOLang主要特征:
        1. 自动垃圾回收
        2. 更丰富的内置类型
        3. 函数多返回值
        4. 错误处理
        5. 匿名函数和闭包
        6. 类型和接口
        7. 并发编程
        8. 反射
        9. 语言交互性
    Golang文件名: 所有的go源码都是以 ".go" 结尾
	Go语言命名:
		Go 的函数, 变量, 常量, 自定义类型, 包(package) 的命名方式遵循以下规则:
			1. 首字符可以是任意的Unicode字符或者下划线
			2. 剩余字符可以是Unicode字符,下划线,数字
			3. 字符长度不限
		Go 语言25个关键字:
			break default func interface select
			case defer go map struct
			chan else goto package switch
			const fallthrough if range type
			continue for import return var
		Go 语言37个保留字:
			Constants: true false iota nil
			Types: int int8 int16 int32 int64
				   uint uint8 uint16 uint32 uint64 uintptr
				   float32 float64 complex128 complex64
				   bool byte rune string error
			Functions: make len cap new append copy close delete
					   complex real imag
					   panic recover
		可见性:
			1. 声明在函数内部, 是函数的本地值, 类似于private
			2. 声明在函数外部, 对当前包可见(包内所有的.go文件都可见), 类似protect
			3. 声明在函数外部且首字母大写是所有包可见的全局值, 类似于public
		Go语言四种声明方式:
			1. var: 声明变量
			2. const: 声明常量
			3. type: 声明类型
			4. func: 声明函数
*/

func main() {

}
