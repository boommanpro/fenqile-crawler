package main

import (
	"fenqile-crawler/config"
	"fenqile-crawler/cos"
	"fenqile-crawler/crawler"
	"fenqile-crawler/serverChan"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
)

var yml string = "config/application.yml"

func main() {
	log.Println("Starting...")

	startTask()
	forYml := config.GetCronForYml(yml)

	// 定义一个cron运行器
	c := cron.New()

	c.AddFunc(forYml.CrawlerCron, startTask)

	c.Start()
	select {}
}

func startTask() {
	tentcentCos := config.GetTencentCosForYml(yml)
	serverChanConfig := config.GetServerChanForYml(yml)
	crawlerConfig := config.GetCrawlerForYml(yml)
	fileMap := crawler.CrawlerData(*crawlerConfig)
	err := cos.UpdateFile(*tentcentCos, fileMap)
	if err != nil {
		fmt.Printf("%s", err)
	}

	sendMsg := ""
	for k, _ := range fileMap {
		sendMsg += fmt.Sprintf("![%s](%s/%s)\n", k, tentcentCos.BaseUrl, k)
	}

	serverChan.SendMessage(*serverChanConfig, "分期乐信息爬虫", sendMsg)
	fmt.Printf("发送信息成功")
}
