package model

import (
	"github.com/jinzhu/gorm"
)

// Comment 评论模型
type Comment struct {
	gorm.Model
	VideoID string
	Info    string
}
