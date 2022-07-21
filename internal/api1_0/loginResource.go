// Package api1_0
/*
@Coding : utf-8
@time : 2022/7/3 21:49
@Author : yizhigopher
@Software : GoLand
*/
package api1_0

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golangOnlineJudge/internal/globals/responseParser"
	"golangOnlineJudge/internal/services"
	"golangOnlineJudge/internal/utils/jwt"
)

type UserApi struct {
}

type loginRequestParser struct {
	UserName string `json:"UserName" form:"UserName" binding:"required"`
	Password string `json:"Password" form:"Password" binding:"required"`
}

func (u *UserApi) Login(c *gin.Context) {
	parser := loginRequestParser{}
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取请求失败", err)
		return
	}

	var userService services.UsersService
	userService.UserName = parser.UserName
	err = userService.Get()
	if err != nil {
		if err.Error() == "record not found" {
			responseParser.JsonNotData(c, "用户不存在", err)
			return
		}
		responseParser.JsonDBError(c, "数据库错误", err)
		return
	}

	cipher := userService.Password
	err = bcrypt.CompareHashAndPassword([]byte(cipher), []byte(parser.Password))
	if err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			responseParser.JsonDataError(c, "密码错误", err)
			return
		}
		responseParser.JsonInternalError(c, "", err)
		return
	}

	token, err := jwt.MakeToken(userService.UserID, userService.IsAdmin)
	if err != nil {
		responseParser.JsonInternalError(c, "生成token失败", err)
		return
	}

	responseParser.JsonOK(c, "登录成功", map[string]interface{}{
		"user":  userService,
		"token": token,
	})
	return
}

type refreshParser struct {
	Token string `json:"Token" binding:"required"`
}

func (u *UserApi) RefreshToken(c *gin.Context) {
	var parser refreshParser
	err := c.ShouldBindJSON(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取Token失败", err)
		return
	}
	token := parser.Token

	token, err = jwt.RefreshToken(token)
	if err != nil {
		responseParser.JsonDataError(c, "token已过期！", err)
		return
	}
	responseParser.JsonOK(c, "更新token成功！", gin.H{
		"token": token,
	})
	return
}
