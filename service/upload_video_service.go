package service

import (
	"giligili/serializer"
	"mime"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

// UploadVideoService 获得上传oss token的服务
type UploadVideoService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建token
func (service *UploadVideoService) Post() serializer.Response {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取扩展名， 需要多拓展名进行验证
	ext := filepath.Ext(service.Filename)

	// 带可选参数的签名直传
	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}

	key := "upload/video/" + uuid.Must(uuid.NewRandom()).String() + ext
	// 签名直传。
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}
	// 不需要视频预览
	//signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	//if err != nil {
	//	return serializer.Response{
	//		Status: 50002,
	//		Msg:    "OSS配置错误",
	//		Error:  err.Error(),
	//	}
	//}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
		},
	}
}
