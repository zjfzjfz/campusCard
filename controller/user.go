package controller

import (
	"campusCard/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u UserController) GetCardInfo(c *gin.Context) {
	id := c.Param("id")
	card, _ := model.GetCard(id)
	ReturnSuccess(c, 200, "success", card)
}

func (u UserController) PutLimit(c *gin.Context) {
	id := c.Param("id")
	limit := c.Param("limit")
	nowLimit, _ := model.ChangeLimit(id, limit)
	ReturnSuccess(c, 200, "success", nowLimit)
}

func (u UserController) Register(c *gin.Context) {
	id := c.DefaultPostForm("id", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")
	username := c.DefaultPostForm("username", "")
	iid := c.DefaultPostForm("iid", "")

	if id == "" || password == "" || confirmPassword == "" || iid == "" || username == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 4004, "两次密码输入不一致")
		return
	}

	user, err := model.GetUserInfoByUserId(id)
	if user.Id != "" {
		ReturnError(c, 4004, "学号已注册")
	}
	_, err = model.AddUser(id, EncryMd5(password), username, iid)
	if err != nil {
		ReturnError(c, 4001, "注册失败")
	}
	ReturnSuccess(c, 200, "注册成功", "")
}

type UserApi struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	id := c.DefaultPostForm("id", "")
	password := c.DefaultPostForm("password", "")
	if id == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	user, _ := model.GetUserInfoByUserId(id)
	if user.Id == "" {
		ReturnError(c, 4004, "用户名或密码不正确")
		return
	}

	if user.Pwd != EncryMd5(password) {
		ReturnError(c, 4004, "用户名或密码不正确")
		return
	}

	session := sessions.Default(c)
	session.Set("login:"+user.Id, user.Id)
	session.Save()
	data := UserApi{Id: user.Id, Username: user.Name}
	ReturnSuccess(c, 0, "登录成功", data)
}
