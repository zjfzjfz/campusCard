package controller

import (
	"campusCard/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

// 激活账号-用户注册
func (u UserController) Register(c *gin.Context) {
	id := c.DefaultPostForm("id", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")
	username := c.DefaultPostForm("username", "")
	iid := c.DefaultPostForm("iid", "")

	if id == "" || password == "" || confirmPassword == "" || iid == "" || username == "" {
		ReturnError(c, 405, "请输入正确的信息")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 406, "两次密码输入不一致")
		return
	}

	user, err := model.GetUserInfoByUserId(id)
	if user.Id == "" {
		user.Id, err = model.AddUser(id, EncryMd5(password), username, iid)
		if err != nil {
			ReturnError(c, 407, "注册失败")
			return
		}
		ReturnSuccess(c, 200, "注册成功", user.Id)
	} else {
		ReturnError(c, 408, "学号已注册")
	}
}

// 用户登录
type UserApi struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	id := c.DefaultPostForm("id", "")
	password := c.DefaultPostForm("password", "")
	if id == "" || password == "" {
		ReturnError(c, 405, "请输入正确的信息")
		return
	}

	user, _ := model.GetUserInfoByUserId(id)
	if user.Id == "" {
		ReturnError(c, 409, "用户名或密码不正确")
		return
	}

	if user.Pwd != EncryMd5(password) {
		ReturnError(c, 409, "用户名或密码不正确")
		return
	}

	session := sessions.Default(c)
	session.Set("login", user.Id)
	session.Save()
	data := UserApi{Id: user.Id, Username: user.Name}
	ReturnSuccess(c, 200, "登录成功", data)
}

// 登出账号
func (u UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)

	// 清除与用户相关的会话数据
	session.Delete("login")

	err := session.Save()
	if err != nil {
		ReturnError(c, 404, "登出失败")
		return
	}
	ReturnSuccess(c, 200, "登出成功", "")
}

// 查询卡信息
func (u UserController) GetCardInfo(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	card, err := model.GetCard(id)
	if err != nil {
		ReturnError(c, 500, err)
		return
	}
	ReturnSuccess(c, 200, "查询成功", card)
}

// 查询欠款信息
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

// 查询交易记录
func (u UserController) GetTradeInfo(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	trades, err := model.GetTrade(id)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}
	ReturnSuccess(c, 200, "交易信息查询成功", trades)
}

// 发起交易
func (u UserController) Trade(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	var transaction model.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		ReturnError(c, 500, err)
		return
	}

	// 使用无缓冲通道来接收结果
	resultChan := make(chan model.TransactionResult)
	defer close(resultChan)

	// 开启 Goroutine 并发处理交易
	go func() {
		newBalance, err := model.InsertTransaction(id, transaction)
		resultChan <- model.TransactionResult{NewBalance: newBalance, Err: err}
	}()
	//等待 Goroutine 完成并处理结果
	result := <-resultChan
	if result.Err != nil {
		ReturnError(c, 500, result.Err)
		return
	}
	ReturnSuccess(c, 200, "success", result.NewBalance)
}

// 更改限额
func (u UserController) PutLimit(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	limit := c.Param("limit")
	nowLimit, err := model.ChangeLimit(id, limit)
	if err != nil {
		ReturnError(c, 500, err)
		return
	}
	ReturnSuccess(c, 200, "更改限额成功", nowLimit)
}

// 挂失请求
func (u UserController) LossPost(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	iid := c.Param("iid")
	user, err := model.GetUserInfoByUserId(id)
	if err != nil {
		ReturnError(c, 500, err)
		return
	}
	if iid == user.IId {
		ReturnSuccess(c, 200, "挂失请求成功", iid)
	} else {
		ReturnError(c, 404, "身份证号信息不符")
	}
}

// 充值
func (u UserController) Charge(c *gin.Context) {
	moneyParam := c.Param("money")
	money, err := strconv.ParseFloat(moneyParam, 64)
	if err != nil {
		ReturnError(c, 400, err)
		return
	}

	session := sessions.Default(c)
	id := session.Get("login").(string)
	newBalance, err := model.ChangeBalance(money, id)
	if err != nil {
		ReturnError(c, 500, err)
		return
	}

	ReturnSuccess(c, 200, "充值成功,当前余额：", newBalance)
}

// 浴室还款
func (u UserController) BathRepayment(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	newBalance, err := model.ChangeBath(id)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}
	ReturnSuccess(c, 200, "还款成功，当前余额：", newBalance)

}

// 图书还款
func (u UserController) LibraryRepayment(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("login").(string)
	newBalance, err := model.ChangeLibrary(id)
	if err != nil {
		// 处理错误，例如记录日志或返回错误响应
		ReturnError(c, 500, err)
		return
	}
	ReturnSuccess(c, 200, "还款成功，当前余额：", newBalance)
}
