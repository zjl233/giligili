package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// oss bucket 相关的操作

// SignedPutURL path 是将要创建的文件在bucket中的具体位置
func SignedPutURL(path string, options []oss.Option) (signedPutURL string) {
	signedPutURL, err := BKT.SignURL(path, oss.HTTPPut, 600, options...)
	if err != nil {
		// todo 更具体的错误处理
		panic(err)
	}
	return
}

// SignedGetURL path 是你想要读取的文件在bucket中的具体位置
func SignedGetURL(path string) (signedGetURL string) {
	signedGetURL, err := BKT.SignURL(path, oss.HTTPGet, 600)
	if err != nil {
		// todo 更具体的错误处理，比如文件不存在
		panic(err)
	}
	return
}
