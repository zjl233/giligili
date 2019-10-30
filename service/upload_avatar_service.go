package service

import (
	"giligili/model"
	"giligili/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"mime"
	"path/filepath"
)

// UploadAvatarService 获得上传oss token的服务
type UploadAvatarService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建token
func (service *UploadAvatarService) Post() serializer.Response {
	// 获取扩展名， 需要多拓展名进行验证
	ext := filepath.Ext(service.Filename)

	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext
	// 签名直传。
	signedPutURL := model.SignedPutURL(key, options)
	// 查看图片
	signedGetURL := model.SignedGetURL(key)

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
