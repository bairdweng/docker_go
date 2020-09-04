package coroutine

import (
	"fmt"
	"time"
)

// CreateChanel 的创建
func CreateChanel() {
	// 通道里面一定用协程。设置缓冲为1
	// ch := make(chan int, 1)
	// go show2(ch)
	// time.Sleep(time.Second * 2)
	// go input(ch)
	// go output(ch)
	// 阻塞主程序退出
	// select {}
}

/// 关闭通道
func show2(ch chan int) {

	num := <-ch
	println("show2读取==========", num)
	// 读取之后关闭通道，关闭通道后不能再写入
	close(ch)
	// 判断通道是否已经关闭
	if _, ok := <-ch; ok == false {
		fmt.Println("通道已经关闭")
		return
	}
	ch <- 20
}
func input(ch chan int) {
	println("写入")
	ch <- 10
	ch <- 11
	for i := 0; i < 100; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
}

func output(ch chan int) {
	for {
		num := <-ch
		time.Sleep(time.Second * 1)
		println("show3读取==========", num)
	}
}
