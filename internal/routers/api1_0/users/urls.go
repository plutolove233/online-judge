// Package users
/*
@Coding : utf-8
@time : 2022/7/3 22:21
@Author : yizhigopher
@Software : GoLand
*/
package users

import (
	"github.com/gin-gonic/gin"
	"golangOnlineJudge/internal/api1_0/users"
)

var (
	api *gin.RouterGroup
)

func InitUsersRouterGroup(engine *gin.RouterGroup) {
	api = engine.Group("users")
	var userApi users.UserApi
	api.POST("register", userApi.Register)
}
