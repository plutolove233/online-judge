package settings

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang-online-judge/internal/middlewares"
	"golang-online-judge/internal/routers"
)

func InitGinEngine() (*gin.Engine, error) {
	gin.SetMode(viper.GetString("system.Mode"))
	engine := gin.Default()
	engine.Static("/static", "static")
	// 加载全局中间件
	engine.Use(middlewares.CorsMiddleware())
	engine.Use(middlewares.LogMiddleware())

	store := cookie.NewStore([]byte(viper.GetString("system.Secret")))

	store.Options(sessions.Options{
		MaxAge: viper.GetInt("system.SessionExpireTime"),
	})
	engine.Use(sessions.Sessions("mySession", store))

	routers.InitStaticRouterGroup(engine)
	routers.InitRouter(engine)

	return engine, nil
}
