package controller

import (
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}
type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (a AdminController) GetCardInfo(c *gin.Context) {
	//cid := c.PostForm("cid")
	//name := c.DefaultPostForm("name", "none")
	//// 将 id 字符串转换为整数
	//ReturnSuccess(c, 0, cid, name)
	search := &Search{}

	err := c.BindJSON(&search)
	if err == nil {
		ReturnSuccess(c, 0, search.Name, search.Cid)
		return
	}
	ReturnError(c, 0, gin.H{"err": err})
}
