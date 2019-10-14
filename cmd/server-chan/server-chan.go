package main

import (
	"fenqile-crawler/config"
	"fenqile-crawler/serverChan"
	"fmt"
)

func main() {

	serverChanConfig := config.GetServerChanForYml("config/application.yml")

	message := serverChan.SendMessage(*serverChanConfig, "boomman", "123456")

	if message != nil {
		fmt.Printf("%s", message)
	}
}
