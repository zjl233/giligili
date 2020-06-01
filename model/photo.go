package model

import (
	"github.com/jinzhu/gorm"
)

// Photo 视频模型
type Photo struct {
	gorm.Model
	AlbumID uint
	Name    string // 其实是阿里云 oss 的 path + uuid name
}

// PhotosURL 封面地址
func (photo *Photo) PhotoURL() string {
	return SignedGetURL(photo.Name)
}
