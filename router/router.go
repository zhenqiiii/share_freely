package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/share_freely/controller"
	"github.com/zhenqiiii/share_freely/pkg/jwt"
)

func SetupRouter(mode string) *gin.Engine {
	// 设置gin的三种模式,默认DebugMode
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 发布模式
	}

	r := gin.Default()

	// user
	userGroup := r.Group("/user")
	{
		// register
		// 注册
		userGroup.POST("/register", controller.RegisterHandler())
		// login
		// 登录
		userGroup.POST("/login", controller.LoginHandler())
	}

	// home
	homeGroup := r.Group("/home")
	{
		// 访问
		homeGroup.GET("/main", controller.HomeHandler())
	}

	// function
	funcGroup := r.Group("/func")
	{
		// 点击换一篇按钮
		funcGroup.GET("/change", controller.ChangePost())
		// 点击上传按钮,跳转页面
		funcGroup.GET("/upload", AuthMiddleware())
		// 确认上传
		funcGroup.POST("/upload", AuthMiddleware(), controller.UploadPost())
		// 查看我的上传
		funcGroup.GET("/myuploads", AuthMiddleware(), controller.MyUploads())
	}

	// 无此路径
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "404",
		})
	})

	return r
}

// ============================MIDDLEWARE==================================//

// 在需要登录的功能中注册该中间件
// 场景：识别到请求中没有token或者token无效
// 响应：后端将响应码及信息发送给前端，前端显示“未登录或登录已过期”
// 或者直接跳转至登录页面

// AuthMiddleware: 基于JWT的认证中间件
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Token存放在Header的Authorization中，使用Bearer开头
		// Authorization: Bearer xxx.xxx.xxx
		// 获取请求Header中的Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" { // 判断Auth是否为空
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}

		// 根据格式，按空格将Authorization分割
		// func SplitN(s, sep string, n int) []string
		// SplitN slices s into substrings separated by sep
		// and returns a slice of the substrings between those separators.
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") { // 判断格式是否正确
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		//parts[1]是tokenstring，接下来进行解析
		claims, err := jwt.ParseToken(parts[1])
		if err != nil { // token invalid
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的userID保存到请求的上下文c中
		c.Set("userid", claims.UserID)
		// 执行后面的处理函数
		// 后续的处理函数可以用过c.Get("userid")来获取当前请求的用户ID信息
		c.Next()
	}
}
