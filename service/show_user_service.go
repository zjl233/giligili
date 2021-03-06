package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowUserService 用户资料的服务
type ShowUserService struct {
}

// Show 用户
func (service *ShowUserService) Show(userName string) serializer.Response {
	var user model.User
	err := model.DB.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "用户不存在",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}
