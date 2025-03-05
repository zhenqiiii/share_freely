package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/share_freely/logic"
	"github.com/zhenqiiii/share_freely/models"
)

// 用户模块的控制台代码：Handler
// Login登录处理函数
func LoginHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1.接收登录参数
		var p models.ParamLogin
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "something went wrong with the param",
				"error": err,
			})
			return
		}
		// 2.调用登录业务逻辑函数
		token, err := logic.Login(&p)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "fail to login",
				"error": err,
			})
			return
		}
		// 3.返回响应及token
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "welcome!",
			"token": token,
		})
	}
}

// Register:注册处理函数
func RegisterHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1.接收注册参数
		var p models.ParamRegister
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "something went wrong with the param",
				"error": err,
			})
			return
		}
		// 2.调用业务逻辑函数
		err = logic.Register(&p)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "fail to register",
				"error": err,
			})
			return
		}
		// 3.返回响应
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Register successfully!",
		})

	}
}
