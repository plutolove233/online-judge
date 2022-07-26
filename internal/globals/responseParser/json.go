// Package responseParser
/*
@Coding : utf-8
@time : 2022/7/3 15:34
@Author : yizhigopher
@Software : GoLand
*/
package responseParser

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golangOnlineJudge/internal/globals/codes"
	"net/http"
)

func JsonOK(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "成功!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": msg,
		"data":    data,
	})
}

func JsonParameterIllegal(c *gin.Context, msg string, err error) {
	if msg == "" {
		msg = "参数非法!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.ParameterIllegal,
		"message": msg,
		"err":     err.Error(),
	})
}

func JsonDataError(c *gin.Context, msg string, err error) {
	if msg == "" {
		msg = "数据错误!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.DataError,
		"message": msg,
		"err":     err.Error(),
	})
}

func JsonNotData(c *gin.Context, msg string, err error) {
	if msg == "" {
		msg = "无数据!"
	}
	if err == nil {
		err = errors.New(msg)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.DataError,
		"message": msg,
		"err":     err.Error(),
	})
}

func JsonInternalError(c *gin.Context, msg string, err error) {
	if msg == "" {
		msg = "系统错误!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.InternalError,
		"message": msg,
		"err":     err.Error(),
	})
	return
}

func JsonDBError(c *gin.Context, msg string, err error) {
	if err.Error() == "record not found" {
		if msg == "" {
			msg = "无数据!"
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.NotData,
			"message": msg,
			"err":     err.Error(),
		})
		return
	}
	if msg == "" {
		msg = "数据库错误!"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.DBError,
		"message": msg,
		"err":     err.Error(),
	})
}

func JsonDataExist(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.DataExist,
		"message": msg,
	})
}

func JsonAccessDenied(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.AccessDenied,
		"message": msg,
	})
}

func JsonLoginError(c *gin.Context, msg string, err error) {
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.LoginError,
			"message": msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    codes.LoginError,
			"message": msg,
			"err":     err,
		})
	}
}

func JsonUnauthorizedUserId(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.UnauthorizedUserId,
		"message": msg,
	})
}

func JsonIncompleteRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.ParameterIllegal,
		"message": msg,
	})
}
