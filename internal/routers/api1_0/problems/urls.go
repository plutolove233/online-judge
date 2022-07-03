// Package problems
/*
@Coding : utf-8
@time : 2022/7/3 22:27
@Author : yizhigopher
@Software : GoLand
*/
package problems

import "github.com/gin-gonic/gin"

var (
	api *gin.RouterGroup
)

func InitProblemsRouterGroup(engine *gin.RouterGroup) {
	api = engine.Group("problems")

}