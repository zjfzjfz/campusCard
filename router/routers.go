package router

import (
	"campusCard/config"
	"campusCard/controller"
	"github.com/gin-contrib/sessions"
	sessionsRedis "github.com/gin-contrib/sessions/redis"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	store, _ := sessionsRedis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mySession", store))
	user := r.Group("/user")
	{
		user.POST("/register", controller.UserController{}.Register)
		user.POST("/login", controller.UserController{}.Login)
		user.GET("/cardinfo/:id", controller.UserController{}.GetCardInfo)
		user.GET("/tradeinfo/:id", controller.UserController{}.GetTradeInfo)
		user.POST("/trade/:id", controller.UserController{}.Trade)
		user.GET("/debtinfo", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller debt")
		})
		user.POST("/limit/:id/:limit", controller.UserController{}.PutLimit)
		user.POST("/bath", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller login")
		})
		user.POST("/library", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller login")
		})
		user.POST("/loss", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller login")
		})
		user.GET("/charge", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller login")
		})
	}
	admin := r.Group("/admin")
	{
		admin.POST("/login", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller login")
		})
		admin.GET("/cardinfo", controller.AdminController{}.GetCardInfo)
		admin.GET("/debtinfo", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller debt")
		})
		admin.PUT("/limit/:limit", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller login")
		})

	}
	return r
}
