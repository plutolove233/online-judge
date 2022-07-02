package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang-online-judge/internal/globals/codes"
	"golang-online-judge/internal/models/ginModels"
	"golang-online-judge/internal/utils/jwt"
	"golang-online-judge/internal/utils/logs"
	"net/http"
)

var log = logs.GetLogger()

func TokenRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token验证
		token := c.Request.Header.Get("Token")
		jwtChaim, err := jwt.VerifyToken(token)
		if err != nil {
			log.Errorln(err)
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.AccessDenied,
				"message": "您的Token已过期！",
			})
			c.Abort()
			return
		}

		var User ginModels.UserModel
		User.UserID = jwtChaim.UserID
		User.IsAdmin = jwtChaim.IsAdmin

		c.Set("user", User)
		c.Next()
	}
}
