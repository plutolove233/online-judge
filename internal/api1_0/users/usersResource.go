// Package users
/*
@Coding : utf-8
@time : 2022/7/3 15:18
@Author : yizhigopher
@Software : GoLand
*/
package users

import (
	"github.com/gin-gonic/gin"
	"golang-online-judge/internal/globals/responseParser"
	"golang-online-judge/internal/services"
	"golang-online-judge/internal/utils/snowflake"
	"golang.org/x/crypto/bcrypt"
)

type UserApi struct {
}

type RegisterRequestParser struct {
	UserName string `json:"UserName" form:"UserName" binding:"required"`
	Password string `json:"Password" form:"Password" binding:"required"`
	Email    string `json:"Email" form:"Email" binding:"required"`
}

func (u *UserApi) Register(c *gin.Context) {
	parser := RegisterRequestParser{}
	err := c.ShouldBind(&parser)
	if err != nil {
		responseParser.JsonParameterIllegal(c, "获取请求参数失败", err)
		return
	}

	var userService services.UsersService
	userService.UserName = parser.UserName

	err = userService.Get()
	if err == nil {
		responseParser.JsonDataExist(c, "该用户名已存在")
		return
	} else if err.Error() != "record not found" {
		responseParser.JsonDBError(c, "", err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(parser.Password), bcrypt.DefaultCost)
	if err != nil {
		responseParser.JsonInternalError(c, "密码加密失败", err)
		return
	}
	userService.UserID = snowflake.GetSnowFlakeID()
	userService.Password = string(hash)

	err = userService.Add()
	if err != nil {
		responseParser.JsonDBError(c, "添加用户信息失败", err)
		return
	}

	responseParser.JsonOK(c, "用户注册成功", userService)
	return
}

func (u *UserApi) SolveProblems(c *gin.Context)  {
	
}