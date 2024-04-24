package controller

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type JsonErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data}
	c.JSON(200, json)
}
func ReturnError(c *gin.Context, code int, msg interface{}) {
	// 尝试将msg断言为error类型
    if errMsg, ok := msg.(error); ok {
        // 如果msg是error类型，获取其字符串表示
        msg = errMsg.Error()
    }
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(404, json)
}

func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
