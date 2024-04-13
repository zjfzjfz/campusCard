package controller

import (
	"campusCard/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

func (u UserController) GetCardInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")
	// 将 id 字符串转换为整数
	id, _ := strconv.Atoi(idStr)
	user, _ := model.GetCard(id)
	ReturnSuccess(c, 0, name, user)
}
