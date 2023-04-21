package main

import (
	"errors"
	"fmt"
)

/*
	error接口:
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
			使用New函数自定义错误:
				==> errors.New("错误信息") ==> 返回error类型对象
			func (e *errorString) Error() string {
				return e.text
			}

		另一个可以生成error类型值的方法是调用fmt包中的Errorf函数:
			package fmt
			import "errors"

			func Errorf(format string, args ...interface{}) error {
				return errors.New(Sprintf(format, args...))
			}
	        使用Errorf函数, 输出格式化错误:
				fmt.Errorf(格式化字符串, 格式化变量)
*/

// 使用New函数自定义错误
func circleArea(radius float64) (float64, error) {
	// 定义函数求圆的面积
	// 定义常量π
	const PI float64 = 3.1415926
	// 判断半径radius是否大于0
	if radius > 0 {
		return PI * radius, nil
	} else {
		return 0.0, errors.New("半径不能小于或等于0！！") // 自定义错误返回
	}
}

// 使用Errorf函数自定义错误
func circleAreaErrorf(radius float64) (float64, error) {
	// 定义函数求圆的面积
	// 定义常量π
	const PI float64 = 3.1415926
	// 判断半径radius是否大于0
	if radius > 0 {
		return PI * radius, nil
	} else {
		return 0.0, fmt.Errorf("圆的半径%0.2f不大于0！！", radius) // 自定义错误返回
	}
}

func main() {
	area, err := circleArea(0)
	if err != nil {
		fmt.Printf("错误信息: %s\n", err)
	} else {
		fmt.Printf("area = %0.2f\n", area)
	}

	areaErrorf, errErrorf := circleAreaErrorf(-20.00)
	if errErrorf != nil {
		fmt.Println(errErrorf)
	} else {
		fmt.Printf("area = %0.2f\n", areaErrorf)
	}
}
