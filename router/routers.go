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
		user.GET("/tradeinfo/:period", controller.UserController{}.GetTradeInfo)
		user.GET("/debtinfo", controller.UserController{}.GetDebtInfo)
		user.GET("/limitinfo", controller.UserController{}.GetLimitInfo)

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

	library := r.Group("/library")
	{
		library.POST("/reserve", controller.LibraryController{}.ReserveSeat)
		library.POST("/quick_reserve", controller.LibraryController{}.QuickReserveSeat)
		library.POST("/scan_reserve", controller.LibraryController{}.ScanReserveSeat)
		library.POST("/check_in", controller.LibraryController{}.CheckIn)
		library.POST("/temporary_leave", controller.LibraryController{}.TemporaryLeave)
		library.POST("/check_out", controller.LibraryController{}.CheckOut)
	}
	return r
}
