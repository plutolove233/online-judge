// Package api1_0
/*
@Coding : utf-8
@time : 2022/7/3 14:58
@Author : yizhigopher
@Software : GoLand
*/
package api1_0

import (
	"github.com/gin-gonic/gin"
	"golang-online-judge/internal/api1_0"
	"golang-online-judge/internal/routers/api1_0/problems"
	"golang-online-judge/internal/routers/api1_0/submits"
	"golang-online-judge/internal/routers/api1_0/users"
)

func InitAPI1_0Router(engine *gin.Engine) {
	api := engine.Group("api_1_0")
	api.Any("version", api1_0.Version)
	userApi := api1_0.UserApi{}
	api.POST("login", userApi.Login)
	api.POST("refreshToken", userApi.RefreshToken)

	users.InitUsersRouterGroup(api)
	problems.InitProblemsRouterGroup(api)
	submits.InitSubmitApiRouterGroup(api)
}
