package coroutine

import (
	"fmt"
	"sync"
	"time"
)

// LockExample 加锁
func LockExample() {
	var mt sync.Mutex
	var wg sync.WaitGroup
	var money = 10000

	// 开启10个协程，每个协程内部 循环1000次，每次循环值+10
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			mt.Lock()
			fmt.Printf("协程 %d 抢到锁\n", index)
			for j := 0; j < 100; j++ {
				money += 10 //  多个协程对 money产生了竞争
			}
			fmt.Printf("协程 %d 准备解锁\n", index)
			mt.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("最终的monet = ", money) // 应该输出20000才正确
}

// CanNotRun 协程不执行
func CanNotRun() {
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(i) // 全部打印11：因为开启协程也会耗时，协程没有准备好，循环已经走完
		}()
	}
	time.Sleep(time.Second)
}

// PrintRandom 打印无规律
func PrintRandom() {
	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i) // 打印无规律数字
		}(i)
	}
	time.Sleep(time.Second)
}
