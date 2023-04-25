package No8_ConcurrentProgramm

/*
	Sync包:
		提供了基本的同步基元, 如互斥锁, 其他常用类型与方法如下:
        1. sync.WaitGroup
            sync.WaitGroup内部维护着一个计数器, 计数器的值可以增加和减少
            方法:
            1.1 (wg *WaitGroup) Add(delta int): 将计数器的值 +delta
            1.2 (wg *WaitGroup) Done(): 将计数器的值 -1
            1.3 (wg *WaitGroup) Wait(): 阻塞程序运行, 直到计数器值 归0
            注意:
                sync.WaitGroup是一个结构体, 传递的时候要传递指针
        2. sync.Once:
            使用场景==>需要确保某些操作在高并发的场景下只执行一次
            只有一个Do方法: func (o *Once) Do(f func()) {}
            注意: 如果要执行的函数f需要传递参数就需要搭配闭包来使用
            定义sync.Once变量:
                var var_name sync.Once
            sync.Once内部包含一个互斥锁和一个布尔值:
                互斥锁: 保证初始化操作是并发安全的
                布尔值: 记录初始化是否完成
                这样设计可以保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次
        3. sync.Map:
            是Go语言中提供的一个并发安全版的map, 该map不需要像内置的Map一样使用make函数初始化即可使用
            定义sync.Map变量:
                var mapName = sync.Map{}
            sync.Map常用方法:
                func (m *Map) Store(key, value any):存储一对Key-value值
                func (m *Map) Load(key any) (value any, ok bool) ==> 根据key获取对应的value, 并判断该key是否存在
                func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool): 如果key对应的value存在, 则返回该value,
																						如果不存在, 存储相应的value
                func (m *Map) Delete(key any): 删除一个key-value键值对
                func (m *Map) Range(f func(key, value any) bool): 循环迭代sync.Map, 效果与for range一样
*/
