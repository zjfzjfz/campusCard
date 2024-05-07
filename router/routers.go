package router

import (
	"campusCard/config"
	"campusCard/controller"
	"github.com/gin-contrib/sessions"
	sessionsRedis "github.com/gin-contrib/sessions/redis"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	store, _ := sessionsRedis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mySession", store))
	user := r.Group("/user")
	{
		user.POST("/register", controller.UserController{}.Register)
		user.POST("/login", controller.UserController{}.Login)
		user.POST("/logout", controller.UserController{}.Logout)

		user.GET("/cardinfo", controller.UserController{}.GetCardInfo)
		user.GET("/tradeinfo", controller.UserController{}.GetTradeInfo)
		user.GET("/debtinfo", controller.UserController{}.GetDebtInfo)

		user.GET("/limit", controller.UserController{}.PutLimit)
		user.PUT("/limit/:limit", controller.UserController{}.PutLimit)

		user.POST("/trade", controller.UserController{}.Trade)
		user.POST("/bath", controller.UserController{}.BathRepayment)
		user.POST("/library", controller.UserController{}.LibraryRepayment)
		user.POST("/loss/:iid", controller.UserController{}.LossPost)
		user.POST("/charge/:money", controller.UserController{}.Charge)
	}

	admin := r.Group("/admin")
	{
		admin.POST("/login", controller.UserController{}.Login)

	}
	return r
}
