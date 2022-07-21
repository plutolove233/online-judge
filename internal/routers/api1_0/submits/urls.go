// Package submits
/*
@Coding : utf-8
@time : 2022/7/13 22:06
@Author : yizhigopher
@Software : GoLand
*/
package submits

import (
	"github.com/gin-gonic/gin"
	"golangOnlineJudge/internal/api1_0/submits"
	"golangOnlineJudge/internal/middlewares"
)

var api *gin.RouterGroup

func InitSubmitApiRouterGroup(engine *gin.RouterGroup) {
	api = engine.Group("submit")
	var submitApi submits.SubmitApi
	api.Use(middlewares.TokenRequire())
	api.POST("submit", submitApi.SubmitProblems)
	api.GET("judge", submitApi.JudgeSubmit)
	api.GET("list", submitApi.GetSubmitRecord)
}
