package router

import (
	"campusCard/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("/login", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller login")
		})
		user.GET("/cardinfo/:id", controller.UserController{}.GetCardInfo)
		user.GET("/tradeinfo", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller trade")
		})
		user.GET("/debtinfo", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "controller debt")
		})
		user.GET("/limit/:id/:limit", controller.UserController{}.PutLimit)
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
