package main

import (
	"fenqile-crawler/config"
	"fenqile-crawler/cos"
	"fenqile-crawler/crawler"
	"fenqile-crawler/serverChan"
	"fmt"
)

func main() {
	fileMap := crawler.CrawlerData()
	tentcentCos := config.GetTencentCosForYml("config/application.yml")
	err := cos.UpdateFile(*tentcentCos, fileMap)
	if err != nil {
		fmt.Printf("%s", err)
	}

	serverChanConfig := config.GetServerChanForYml("config/application.yml")

	sendMsg := ""
	for k, _ := range fileMap {
		sendMsg += fmt.Sprintf("![%s](%s/%s)\n", k, tentcentCos.BaseUrl, k)
	}

	serverChan.SendMessage(*serverChanConfig, "分期乐信息爬虫", sendMsg)
}
