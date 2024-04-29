package main_test

import (
	"bytes"
	"campusCard/controller"
	"encoding/json"
	"campusCard/model"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ExpectedResponse struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}



func TestTrade(t *testing.T) {
	// 创建一个带有随机密钥的 cookie 存储
    store := cookie.NewStore([]byte("secret"))
	
	// 设置 Gin 引擎为测试模式
	gin.SetMode(gin.TestMode)

	requestBody := model.Transaction{
        TType     :0,
		TLocation :"yixing",
		TTime     :"2024/04/29",
		TAmount   :2,
    }
	requestBodyBytes, err := json.Marshal(requestBody)
    if err != nil {
        t.Fatalf("failed to marshal request body: %v", err)
    }

	// 创建一个 Gin 引擎
	r := gin.New()
	r.Use(sessions.Sessions("mysession", store))

	// 注册路由
	r.POST("/trade", func(c *gin.Context) {
        session := sessions.Default(c)
        session.Set("login", "235") // 假设用户已登录，将用户 ID 设置为 "123"
        session.Save()
        // 这里调用你的处理函数
        controller.UserController{}.Trade(c)
    })

	// 创建一个等待组，用于等待所有并发请求完成
	var wg sync.WaitGroup
	wg.Add(10000) // 设置等待组的计数器为 100

	// 并发发送 100 个请求
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done() // 减少等待组的计数器

			// 构造一个模拟的 HTTP 请求
			req, _ := http.NewRequest("POST", "/trade", bytes.NewBuffer(requestBodyBytes))
			req.Header.Set("Content-Type", "application/json")
			// 创建一个 ResponseRecorder 来记录响应
			rr := httptest.NewRecorder()

			// 处理请求
			r.ServeHTTP(rr, req)

			// 检查状态码是否为 200
			assert.Equal(t, http.StatusOK, rr.Code, "Status code not as expected")

			// 检查返回的 JSON 数据是否包含 "success" 字段
			expectedResponse := ExpectedResponse{
				Code: 200,
				Msg:  "success",
				Data: nil, // 在这里可以使用任何预期的数据
			}
			var actualResponse ExpectedResponse
			err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)
			if err != nil {
				//t.Fatalf("failed to unmarshal actual response: %v", err)
				err=nil
			}

			assert.Equal(t, expectedResponse.Code, actualResponse.Code, "Code not as expected")
			assert.Equal(t, expectedResponse.Msg, actualResponse.Msg, "Msg not as expected")
			// 检查 Data 字段是否与预期相符

		}()
	}

	// 等待所有并发请求完成
	wg.Wait()
}

