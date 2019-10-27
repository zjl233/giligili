package serializer

import (
	"giligili/model"
	"strconv"
)

// Comment 评论序列化器
type Comment struct {
	ID        uint      `json:"id"`
	Info      string    `json:"info"`
	CreatedAt int64     `json:"created_at"`
	Replies   []Comment `json:"replies"`
}

// BuildComment 序列化评论
func BuildComment(item model.Comment) Comment {
	return Comment{
		ID:        item.ID,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
		Replies:   BuildReplies(strconv.Itoa(int(item.ID))),
	}
}

// BuildComments 序列化评论列表
func BuildComments(items []model.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}

func BuildReplies(parentID string) []Comment {
	var replies []model.Comment
	model.DB.Where("parent_id = ?", parentID).Find(&replies)
	return BuildComments(replies)
}
