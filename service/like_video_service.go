package service

import (
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"github.com/go-redis/redis"
	"time"
)

// LikeVideoService 投稿详情的服务
type LikeVideoService struct {
	VideoID string `json:"video_id"`
}

// Like 视频
func (service *LikeVideoService) Like(user *model.User) serializer.Response {
	// add [timestamp, user_id] pair to `video liked` zset
	// redis-cli> zadd video:10:liked 1573352457 1
	cache.RedisClient.ZAdd(cache.VideoLikedKey(service.VideoID), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: user.ID,
	})

	// todo add [timestamp, video_id] pair to `user like video` zset
	return serializer.Response{
		Status: 0,
		Msg:    "视频 ❤ +1 😎",
	}
}
