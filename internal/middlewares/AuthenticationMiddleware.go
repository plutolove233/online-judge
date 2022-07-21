package middlewares

import (
	"github.com/gin-gonic/gin"
	"golangOnlineJudge/internal/globals/codes"
	"golangOnlineJudge/internal/models/ginModels"
	"net/http"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//查询账号信息
		temp, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.InternalError,
				"message": "用户信息获取错误！",
			})
			c.Abort()
			return
		}
		user := temp.(ginModels.UserModel)
		// 验证权限
		if ok := user.VerifyAdminRole(); !ok {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.AccessDenied,
				"message": "您无权访问！",
			})
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
