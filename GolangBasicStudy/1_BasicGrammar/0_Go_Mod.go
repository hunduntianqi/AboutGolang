/*
	开启go modules功能:
		go env -w GO111MODULE=on
	初始化Go module:
		go mod init project_name
	go module介绍:
		go module 是 go 官方自带的go依赖管理库, 在1.13版本正式推荐使用
		go module 可以将某个项目(文件夹)下的所有依赖整理成一个 go.mod 文件, 里面写入了依赖的版本等, 使用 go module 之后我们可不用将代码放置在src下
		使用 go module 管理依赖后会在项目根目录下生成两个文件 go.mod（会记录当前项目的所依赖）和go.sum(记录每个依赖库的版本和哈希值)
	GO111MODULE:
		GO111MODULE是 go modules 功能的开关
		1. GO111MODULE=off，无模块支持，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找
		2. GO111MODULE=on，模块支持，go命令行会使用modules，而一点也不会去GOPATH目录下查找
		3. GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能, 这种情况下可以分为两种情形：
			（1）当前目录在 GOPATH/src 之外且该目录包含go.mod文件，开启模块支持
			（2）当前文件在包含 go.mod 文件的目录下面
	Go mod使用方法:
		1. 初始化模块 ==> Go mod init project_name
		2. 依赖关系处理 ==> Go mod tidy
		3. 将依赖包复制到项目的 vendor 目录
		4. 显示依赖关系 ==> Go list -m all
		5. 显示详细的依赖关系 ==> Go list -m -json all
		6. 下载依赖 ==> Go mod download [path@version]
*/

package main
