package main

import (
    "fmt"
    "net/http"

    "campusCard/routes"
)

func main() {
    // 初始化路由
    router := routes.SetupRoutes()

    // 启动服务器
    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", router)
}