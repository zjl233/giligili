package serializer

import "giligili/model"

// Comment 评论序列化器
type Comment struct {
	ID        uint   `json:"id"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildComment 序列化评论
func BuildComment(item model.Comment) Comment {
	return Comment{
		ID:        item.ID,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
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
