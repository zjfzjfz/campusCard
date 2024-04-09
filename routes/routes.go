package routes

import (
    "github.com/gin-gonic/gin"

	"campusCard/api"
)

// SetupRoutes 设置路由
func SetupRoutes() *gin.Engine {
    router := gin.Default()

    // 定义路由
    router.GET("/", api.HomeHandler)

    return router
}
