package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"io/ioutil"
	"log"
	"time"
)

//logic:

//1.设置浏览器参数

//2.目标网址

//3.截图保存

//4.循环获取所有   3.4.5.6 3是当前页面  只需要4.5.6

//邮件逻辑

//每天晚上10点和早上9点发送邮件

//正文为3张图片即可

func main() {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	//图片加载设置
	prefs := map[string]interface{}{
		// 2是无图
		//"profile.managed_default_content_settings.images": 2,
		"devtools.responsive.touchSimulation.enabled": true,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: prefs,
		Path:  "",
		Args: []string{
			"--headless",
			//"--auto-open-devtools-for-tabs",
		},
		MobileEmulation: &chrome.MobileEmulation{
			//DeviceName: "iPhone 6",
			DeviceMetrics: &chrome.DeviceMetrics{
				Width:  800,
				Height: 10000,
			},
		},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	service, err := selenium.NewChromeDriverService("D:/go_project/crawler_template/opt/google/chrome/chromedriver.exe", 9515, opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}

	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}

	result := crawlerFenqile(webDriver)
	for i, v := range result {
		err = ioutil.WriteFile(fmt.Sprintf("test%d.jpg", i), v, 0644)
	}
	if err != nil {
		panic(err)

	}

	defer service.Stop()
}

func crawlerFenqile(webDriver selenium.WebDriver) [3][]byte {
	var err error

	err = webDriver.Get("https://hui.m.fenqile.com/")
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	var result [3][]byte

	for i := 0; i < 3; i++ {
		webElement, _ := webDriver.FindElement(selenium.ByCSSSelector, fmt.Sprintf("body > div > div.weex-ct.weex-div.list-nav.scroller-border > main > article > div:nth-child(%d)", i+4))
		webElement.Click()
		bytes := doScreenshot(webDriver)
		result[i] = bytes
	}
	return result

}

func doScreenshot(webDriver selenium.WebDriver) []byte {
	time.Sleep(time.Second * 3)
	bytes, err := webDriver.Screenshot()
	if err != nil {
		return []byte("")
	}
	return bytes
}
