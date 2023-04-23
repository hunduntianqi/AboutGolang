package main

import "fmt"

/*
	Golang异常处理:
        Go语言中没有结构化异常, 使用panic抛出异常, recover捕获异常
        异常的使用场景描述:
            Go中可以抛出一个panic的异常, 然后在defer中通过recover捕获这个异常, 然后正常处理
        panic:
            1. 内置函数 ==> func panic(v interface{}) ==> 空接口类型, 可以抛出任何类型对象
            2. 假如函数F中书写了panic语句, 会终止其后要执行的代码, 在panic所在函数F内如果存在要执行的defer函数列表,
               按照defer的逆序执行
            3. 返回函数F的调用者G, 在G中, 调用函数F之后的代码不会被执行, 假如函数G中存在要执行的defer函数列表, 按照defer的逆序执行
            4. 直到goroutine整个退出, 并报告错误
        recover:
            1. 内置函数 ==> func recover(v interface{}) ==> 空接口类型, 可以抛出任何类型对象
            2. 用来控制一个goroutine的panicking行为, 捕获panic, 从而影响应用的行为
            3. 一般的调用建议:
                a). 在defer函数中, 通过recover来终止一个goroutine的panicking过程, 从而恢复正常代码的执行
                b). 可以获取通过panic传递的error
		基本使用格式:
			func func_name(param list) (return list) {
				// 定义延迟调用捕获异常
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err) // 当捕获到异常时, 会输出异常提示
					}
				}()
			}
        注意:
            1.利用recover处理panic指令, defer 必须放在 panic 之前定义, 另外 recover 只有在 defer 调用的函数中才有效
              否则当panic时, recover无法捕获到panic, 无法防止panic扩散
            2.recover 处理异常后, 逻辑并不会恢复到 panic 那个点去, 函数跑到 defer 之后的那个点
            3.多个 defer 会形成 defer 栈, 后定义的 defer 语句会被最先调用
		使用异常处理注意事项:
			1. 向已关闭的通道内发送数据会引发panic
			2. 延迟调用中引发的错误, 可被后续延迟调用捕获, 但仅最后一个错误可被捕获
			3. 捕获函数 recover 只有在延迟调用内直接调用才会终止错误, 否则总是返回 nil, 任何未捕获的错误都会沿调用堆栈向外传递
				即: 只有在defer修饰的匿名函数内调用recover()函数才可捕获异常
			4. 在一个函数中定义异常捕获语句, 在另外一个函数延迟调用该函数, 也可实现异常捕获; 该方法是将匿名函数替换为一个只有
				异常捕获语句的外部函数, 进行调用
			5. 如果需要保护代码块, 可将代码块重构成匿名函数, 确保后续代码被执行
		error接口使用:
				Go语言引入了一个关于错误处理的标准模式, 即error接口, 是Go语言内建的接口类型
			该接口定义如下:
				type error interface{
					Error() string
				}
			Go语言标准库代码包errors为用户提供如下方法:
				package errors
				type errorString struct {
					text string
				}

				func New(text string) error {
					return &errorString(text)
				}

				func (e *errorString) Error() string {
					return e.text
				}
			另一个可以生成error类型值的方法是调用fmt包中的Errorf函数:
				package fmt
				import "errors"
				func Errorf(format string, args ...interface{}) error {
					return errors.New(Sprintf(format, args...))
				}
	一般情况下: 会导致不可修复性错误的使用panic, 其他情况下使用error
*/

func main() {
	panicTest1()
	panicTest2()
	panicTest3()
	panicTest4()
	panicTest5(2, 1)
}

// 向已关闭的通道内发送数据会引发panic
func panicTest1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	
	// 定义管道类型数据
	var ch chan int = make(chan int, 10)
	// 关闭管道
	close(ch)
	// 关闭管道后, 向管道添加数据, 会引发panic异常
	ch <- 5
}

// 延迟调用中引发的错误, 可被后续延迟调用捕获, 但仅最后一个错误可被捕获
func panicTest2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	
	defer func() {
		panic("defer panic") // 在捕获正常代码异常后, 延迟调用出现异常, 因此捕获该异常, 替代之前的异常, 之前的异常相当于被丢弃
	}()
	
	panic("test panic") // 延迟调用在其他代码之后, 因此先捕获此异常
}

// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误, 否则总是返回 nil, 任何未捕获的错误都会沿调用堆栈向外传递
func panicTest3() {
	defer func() {
		fmt.Println(recover()) // 有效捕获异常
	}()
	defer recover()              // 捕获异常无效
	defer fmt.Println(recover()) // 捕获异常无效
	defer func() {
		func() {
			fmt.Println("defer inner")
			recover() // 捕获异常无效
		}()
	}()
	
	panic("test panic")
}

// 在一个函数中定义异常捕获语句, 在另外一个函数延迟调用该函数, 也可实现异常捕获
// 定义函数, 函数中仅定义异常捕获语句
func except() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

// 定义函数, 在函数中延迟调用except函数, 起到匿名函数的效果
func panicTest4() {
	// 延迟调用except函数
	defer except()
	panic("panic test")
}

// 如果需要保护代码块, 可将代码块重构成匿名函数, 确保后续代码被执行
func panicTest5(x, y int) {
	var z int = 3
	// 将要执行的代码打包为匿名函数
	func() {
		defer func() { // 3. 第三步执行
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic") // 1. 先执行
		z = x / y           // 2. 第二步执行
		return
	}()
	fmt.Printf("x / y = %d\n", z)
}
