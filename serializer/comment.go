package serializer

import "giligili/model"

// Comment 评论序列化器
type Comment struct {
	ID      uint   `json:"id"`
	VideoId uint   `json:"video_id"`
	Info    string `json:"info"`
}

// BuildComment 序列化评论
func BuildComment(item model.Comment) Comment {
	return Comment{
		ID:      item.ID,
		VideoId: item.VideoId,
		Info:    item.Info,
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
