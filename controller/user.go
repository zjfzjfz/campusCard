package controller

import (
	"campusCard/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
}

func (u UserController) GetCardInfo(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	card, err := model.GetCard(id)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}
	ReturnSuccess(c, 200, "success", card)
}

func (u UserController) GetTradeInfo(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	trades, err := model.GetTrade(id)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}
	ReturnSuccess(c, 200, "success", trades)
}

func (u UserController) Trade(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	var transaction model.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 使用无缓冲通道来接收结果
	resultChan := make(chan model.TransactionResult)
	defer close(resultChan)

	// 开启 Goroutine 并发处理交易
	go func() {
		newBalance, err:= model.InsertTransaction(id, transaction)
		resultChan <- model.TransactionResult{NewBalance: newBalance, Err: err}
	}()
	 //等待 Goroutine 完成并处理结果
	result := <-resultChan
	if result.Err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, result.Err)
		return
	}
	ReturnSuccess(c, 200, "success", result.NewBalance)
}

func (u UserController) PutLimit(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	limit := c.Param("limit")
	nowLimit, err := model.ChangeLimit(id, limit)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}
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
	if user.Id == "" {
		_, err = model.AddUser(id, EncryMd5(password), username, iid)
		if err != nil {
			ReturnError(c, 4001, "注册失败")
			return
		}
		ReturnSuccess(c, 200, "注册成功", "")
	} else {
		ReturnError(c, 4004, "学号已注册")
	}
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
	session.Set("login", user.Id)
	session.Save()
	data := UserApi{Id: user.Id, Username: user.Name}
	ReturnSuccess(c, 0, "登录成功", data)
}
func (u UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)

	// 清除与用户相关的会话数据
	session.Delete("login")

	err := session.Save()
	if err != nil {
		// Handle error, if any
		// You may choose to log the error or handle it in some other way
		// For example:
		c.String(http.StatusInternalServerError, "Error occurred while logging out")
		return
	}

}

func (u UserController) GetDebtInfo(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	debts, err := model.GetDebt(id)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}
	ReturnSuccess(c, 200, "欠款查询成功", debts)
}

func (u UserController) LossPost(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	iid := c.Param("iid")
	user, err := model.GetUserInfoByUserId(id)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, "error retrieving user information")
		return
	}
	if iid == user.IId {
		ReturnSuccess(c, 200, "success", iid)
	} else {
		ReturnError(c, 404, "error")
	}
}

func (u UserController) Charge(c *gin.Context) {
	moneyParam := c.Param("money")
	money, err := strconv.ParseFloat(moneyParam, 64)
	if err != nil {
		ReturnError(c, 400, err)
		return
	}

	session := sessions.Default(c)
	id := session.Get("login").(string)
	err = model.ChangeBalance(money, id)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}

	ReturnSuccess(c, 200, "success", "")
}

func (u UserController) BathRepayment(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	err := model.ChangeDebt(id)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}

	ReturnSuccess(c, 200, "success", "")

}

func (u UserController) LibraryRepayment(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	err := model.ChangeLibrary(id)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}

	ReturnSuccess(c, 200, "success", "")
}
