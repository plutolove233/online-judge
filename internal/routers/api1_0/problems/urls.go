// Package problems
/*
@Coding : utf-8
@time : 2022/7/3 22:27
@Author : yizhigopher
@Software : GoLand
*/
package problems

import (
	"github.com/gin-gonic/gin"
	"golang-online-judge/internal/api1_0/problems"
	"golang-online-judge/internal/middlewares"
)

var (
	api *gin.RouterGroup
)

func InitProblemsRouterGroup(engine *gin.RouterGroup) {
	api = engine.Group("problems")
	problemApi := problems.ProblemApi{}
	api.Use(middlewares.TokenRequire())
	api.POST("createProblem", middlewares.AuthenticationMiddleware(), problemApi.UploadNewProblem)
	api.POST("uploadTestCases", middlewares.AuthenticationMiddleware(), problemApi.UploadProblemTestCases)
	api.GET("getProblemList", problemApi.GetProblemList)
}
