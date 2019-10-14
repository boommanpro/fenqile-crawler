package cos

import (
	"bytes"
	"context"
	"fenqile-crawler/config"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func UpdateFile(tentcentCos config.TentcentCos, fileMap map[string][]byte) error {
	var err error
	u, _ := url.Parse(tentcentCos.BaseUrl)
	b := &cos.BaseURL{BucketURL: u}

	tencentTransport := &cos.AuthorizationTransport{
		SecretID:  tentcentCos.SecretId,
		SecretKey: tentcentCos.SecretKey,
	}
	if tentcentCos.Debug {
		tencentTransport.Transport = &debug.DebugRequestTransport{
			RequestHeader: true,
			// Notice when put a large file and set need the request body, might happend out of memory error.
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   true,
		}
	}
	c := cos.NewClient(b, &http.Client{
		Transport: tencentTransport,
	})

	for k, v := range fileMap {
		_, err = c.Object.Put(context.Background(), k, bytes.NewReader(v), nil)
		if err != nil {
			return err
		}
	}
	return nil

}
