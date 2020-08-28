package controllers

import (
	"fmt"
	"runtime"
	"time"

	"com.miaoyou.server/helper"
	"github.com/gin-gonic/gin"
)

// Start 开始
func Start(c *gin.Context) {
	// devExample1()
	devExample3()
	c.JSON(200, helper.Successful(""))

}

// 协程没有准备好
func devExample1() {
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(i) // 全部打印11：因为开启协程也会耗时，协程没有准备好，循环已经走完
		}()
	}
	time.Sleep(time.Second)
}

// 打印无规律
func devExample2() {
	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i) // 打印无规律数字
		}(i)
	}
	time.Sleep(time.Second)
}

// 终止当前协程
func devExample3() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println("defer ", i)
		go func(i int) {
			if i == 3 {
				runtime.Goexit()
			}
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}
