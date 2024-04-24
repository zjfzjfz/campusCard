package main

import (
	"campusCard/dao"
	"campusCard/router"
	"fmt"
)

func main() {
	r := router.Router()

	fmt.Printf("Server running on port 8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}

	defer dao.Db.Close()
}
