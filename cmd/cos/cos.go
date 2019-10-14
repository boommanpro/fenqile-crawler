package main

import (
	"fenqile-crawler/config"
	"fenqile-crawler/cos"
	"fmt"
	"io/ioutil"
)

func main() {

	tentcentCos := config.GetTencentCosForYml("config/application.yml")

	test0, e := ioutil.ReadFile("opt/test0.jpg")
	if e != nil {
		panic(e)
	}
	test1, e := ioutil.ReadFile("opt/test1.jpg")
	if e != nil {
		panic(e)
	}
	test2, e := ioutil.ReadFile("opt/test2.jpg")
	if e != nil {
		panic(e)
	}

	fileMap := map[string][]byte{
		"20191014/test0.jpg": test0,
		"20191014/test1.jpg": test1,
		"20191014/test2.jpg": test2,
	}

	e = cos.UpdateFile(*tentcentCos, fileMap)

	if e != nil {
		fmt.Printf("%s", e)
	}
}
