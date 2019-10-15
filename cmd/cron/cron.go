package main

import (
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	log.Println("Starting...")

	// 定义一个cron运行器
	c := cron.New()
	// 定时5秒，每5秒执行print5
	c.AddFunc("*/5 * * * * *", print5)
	// 定时15秒，每5秒执行print5
	c.AddFunc("*/15 * * * * *", print15)

	// 开始
	c.Start()
	select {}
}

func print5() {
	log.Println("Run 5s cron")
}

func print10() {
	log.Println("Run 10s cron")
}

func print15() {
	log.Println("Run 15s cron")
}
