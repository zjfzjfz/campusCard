package controller

import (
	"campusCard/model"
	"github.com/gin-gonic/gin"
	_ "strconv"
)

type UserController struct {
}

func (u UserController) GetCardInfo(c *gin.Context) {
	id := c.Param("id")
	
	// 将 id 字符串转换为整数
	card, _ := model.GetCard(id)
	ReturnSuccess(c, 200, "success", card)
}
