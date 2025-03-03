package controller

import "github.com/gin-gonic/gin"

// 用户模块的控制台代码：Handler
// Login登录处理函数
func LoginHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1.接收登录参数
		// 2.调用登录业务逻辑函数
		// 3.返回响应及token
	}
}

// Register:注册处理函数
func RegisterHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1.接收注册参数
		// 2.调用业务逻辑函数
		// 3.返回响应
	}
}
