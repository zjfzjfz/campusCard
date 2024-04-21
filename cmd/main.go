package main

import (
	"campusCard/router"
	"campusCard/dao"
	"fmt"
)

func main() {
	// 初始化路由
	r := router.Router()

	// 启动服务器
	fmt.Printf("Server running on port 8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}

	defer dao.Db.Close()
}
