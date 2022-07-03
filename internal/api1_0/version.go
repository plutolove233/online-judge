// Package api1_0
/*
@Coding : utf-8
@time : 2022/7/3 14:54
@Author : yizhigopher
@Software : GoLand
*/
package api1_0

import (
	"github.com/gin-gonic/gin"
	"golang-online-judge/internal/globals/codes"
)

func Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    codes.OK,
		"message": "api.1.0",
	})
}
