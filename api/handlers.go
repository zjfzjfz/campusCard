package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// HomeHandler 处理根路径请求
func HomeHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to my API!",
    })
}