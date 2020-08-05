package helper

import (
	"fmt"

	"github.com/robfig/cron"
	// cron "github.com/robfig/cron/v3"
)

// Corestart 开始
func Corestart() {

	fmt.Print("执行写成")
	cronTab := cron.New()

	cronTab.AddFunc("* * * * *", hello)
	cronTab.Start()
	// select {}
}

func hello() {
	fmt.Println("hello world")
}
