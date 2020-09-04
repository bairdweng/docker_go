package coroutine

import (
	"fmt"
	"sync"
	"time"
)

// CondLock 条件变量
func CondLock() {
	cond := sync.NewCond(&sync.Mutex{})
	condition := false
	// 开启一个新的协程，修改变量 condition

	//如：摄像头，麦克风这类共享资源同一时间内只允许一个人，你想用时发现已经有人在使用，你的选择可以是
	go func() {

		time.Sleep(time.Second * 1)
		cond.L.Lock()

		condition = true // 状态变更，发送通知
		cond.Signal()    // 发信号
		println("等待2s")
		time.Sleep(time.Second * 2)
		cond.L.Unlock()
	}()

	// main协程 是被通知的对象，等待通知
	cond.L.Lock()
	for !condition {
		cond.Wait() // 内部释放了锁（释放后，子协程能拿到锁），并等待通知（消息）
		fmt.Println("获取到了消息")
	}
	cond.L.Unlock() // 接到通知后，会被再次锁住，所以需要在需要的场合释放

	fmt.Println("运行结束")
}

// Person 人们
type Person struct {
	Name string
	Age  int
}

// Grown 长大一岁啦
func (p *Person) Grown() {
	p.Age++
}

// SyncOnce 只执行一次
func SyncOnce() {
	var once sync.Once
	var wg sync.WaitGroup

	p := &Person{
		"比尔",
		0,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				p.Grown()
			})
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("年龄是：", p.Age)
}
