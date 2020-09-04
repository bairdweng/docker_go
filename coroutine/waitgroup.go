package coroutine

import (
	"sync"
)

//例子1 启用10个线程，但所有协程完成的时候发出通知。
var wg sync.WaitGroup // 声明一个等待组
// 还是得加锁
var lock sync.Mutex

// WaitGo 等待组
func WaitGo() {
	println("example1")
	count := 0
	// 缓冲区跟写入数不一致将发生阻塞
	ch := make(chan int, 2)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			lock.Lock()
			count++
			lock.Unlock()
			ch <- 10
			wg.Done()
		}(i)
	}
	// 一直等着,发生了通道阻塞，解决阻塞的方法是用新的协程去读取。或者增加缓冲通道
	go read(ch)
	select {}
}

// 这里发生channel阻塞
func read(ch chan int) {
	println("读出来可以吗=====", <-ch)
	wg.Wait()
	println("hello=====")
}
