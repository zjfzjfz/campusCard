package main

import (
	"campusCard/dao"
	"campusCard/cache"
	"campusCard/router"
	"campusCard/logger"
	"campusCard/scheduledTask"
	"fmt"
)

func main() {
	r := router.Router()

	fmt.Printf("Server running on port 8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}

	defer logger.FileClose()
	defer dao.Db.Close()
	defer cache.Rdb.Close()
	defer scheduledTask.C.Stop()
}
