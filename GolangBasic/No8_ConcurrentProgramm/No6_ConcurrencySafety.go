package No8_ConcurrentProgramm

/*
	并发安全:
        在Go语言中, 可能会存在多个goroutine同时操作一个资源, 这种情况会发生资源竞争问题, 导致可能与预期不符
    互斥锁:
        是一种常用的控制共享资源访问的方法, 可以保证同一时间只有一个goroutine可以访问共享资源
        Go语言使用sync包的 Mutex 类 来实现互斥锁, 常用方法如下:
			1. func (m *Mutex) Lock(): 加互斥锁
			2. func (m *Mutex) Unlock(): 解锁
            互斥锁使用步骤:
                1. var lock sync.Mutex ==> 定义一个互斥锁对象
                2. lock.Lock() ==> 在需要加锁的公共资源代码前调用方法加锁
                3. lock.Unlock() ==> 对公共资源操作完毕后, 在公共资源代码后解锁
                当有goroutine在操作被上锁的公共资源时, 其他Goroutine只有等待解锁后, 才可以操作公共资源,
                当有多个Goroutine等待一个锁时, 解锁后会随机唤醒一个goroutine操作公共代码
	    读写互斥锁:
        读写锁:
            读锁: 当一个goroutine获取读锁时, 其他goroutine如果获取读锁, 会继续获得锁; 如果获取写锁, 则需要等待读锁释放
            写锁: 当一个goroutine获取写锁之后, 其他的goroutine无论是获取读锁还是写锁都会等待
        读写锁定义:
            Go语言使用sync包的 RWMutex 类实现
            var rwlock sync.RWMutex ==> 定义一个读写锁
            rwlock.Lock() ==> 加写锁
            // 需要上锁的代码
            rwlock.Unlock() ==> 解写锁
            rwlock.RLock() ==> 加读锁
            // 需要上锁代码
            rwlock.RUnlock() ==> 解读锁
        注意:
            读写锁适合读多写少的场景, 如果读写操作差别不大, 就无法发挥读写锁的优势
		读写锁常用方法:
			1. func (rw *RWMutex) Lock(): 加写锁, 效果与互斥锁一致
			2. func (rw *RWMutex) Unlock(): 解写锁, 效果与互斥锁一致
			3. func (rw *RWMutex) RLock(): 加读锁, 禁止其他线程写入, 不禁止读取
			4. func (rw *RWMutex) RUnlock(): 解读锁
*/
