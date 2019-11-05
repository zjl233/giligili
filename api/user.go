package api

import (
	"giligili/serializer"
	"giligili/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var serv service.UserRegisterService
	if err := c.ShouldBind(&serv); err == nil {
		if user, err := serv.Register(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var serv service.UserLoginService
	if err := c.ShouldBind(&serv); err == nil {
		if user, err := serv.Login(); err != nil {
			c.JSON(200, err)
		} else {
			// 设置Session
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.ID)
			s.Save()

			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "登出成功",
	})
}

func UserShow(c *gin.Context) {
	serv := service.UserShowService{}
	res := serv.Show(c.Param("user_name"))
	c.JSON(200, res)
}
