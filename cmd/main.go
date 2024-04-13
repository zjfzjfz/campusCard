package main

import (
	"campusCard/router"
	"fmt"
)

func main() {
	// 初始化路由
	r := router.Router()

	// 启动服务器
	fmt.Println("Server is running on port 8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
