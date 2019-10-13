package main

import (
	"context"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func main() {
	var err error
	u, _ := url.Parse("https://fenqile-1252017882.cos.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "xxx",
			SecretKey: "xx",
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	// Case3 put object by local file path
	_, err = c.Object.PutFromFile(context.Background(), "20191013/test4.jpg", "test3.jpg", nil)
	if err != nil {
		panic(err)
	}

}
