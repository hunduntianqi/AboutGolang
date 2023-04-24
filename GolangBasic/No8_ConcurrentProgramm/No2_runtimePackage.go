package No8_ConcurrentProgramm

/*
	runtime包:
		runtime包提供和go运行时环境的互操作
		可用函数:
			1. func NumCPU() int: 返回本地机器的逻辑CPU个数
			2. func GOMAXPROCS(n int) int: 设置可同时执行任务的最大CPU数, 并返回先前的设置
			3. func GC(): 执行一次垃圾回收
			4. func NumGoroutine() int: 获取当前正在运行的 goroutine 数
			5. func Goexit(): 终止当前 goroutine, 并在终止前执行所有 defer 函数
			6. func Gosched(): 暂时出让CPU执行权限, 让其他任务可以获取CPU执行权限, 但不意味着需要等待其他任务执行完毕, 在之后的某个时刻
								可能再次抢夺CPU执行权限
*/
