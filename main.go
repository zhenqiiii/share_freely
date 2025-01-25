package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	ID       int    `json:"id" form:"id"` //注册时由数据库指定：自增
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func main() {
	//mysql database
	dsn := "root:root@tcp(127.0.0.1:3306)/share_freely?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//模型迁移
	db.AutoMigrate(&User{})

	//route handlers
	r := gin.Default()
	//user:用户模块
	userGroup := r.Group("/user")
	{
		//register:注册
		//访问register界面
		userGroup.GET("/register", func(c *gin.Context) {
			//返回模板 register.html
			//c.HTML(http.StatusOK, "register.html", nil)
		})
		//进行注册操作
		userGroup.POST("/register", func(c *gin.Context) {
			//form表单数据:参数绑定
			var user User
			if err := c.ShouldBind(&user); err != nil { //出错
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 2,
					"msg":  err.Error(),
				})
				return
			}
			result := db.Where("username = ?", user.Username).First(&user)
			//找到同名已存在用户
			if result.RowsAffected != 0 {
				c.JSON(http.StatusOK, gin.H{ //响应
					"code": 0,
					"msg":  "Fail to register, occupied username",
				})
				return
			}
			//未找到,进行注册
			result = db.Create(&user)
			if result.Error != nil {
				panic(result.Error)
			}
			c.JSON(http.StatusOK, gin.H{ //响应
				"code": 1,
				"msg":  "Registered successfully",
			})

			// 跳转至登录页:重定向

		})

		// login: 登录
		// 访问登录界面
		userGroup.GET("/login", func(c *gin.Context) {
			//返回login.html
			//c.HTML(http.StatusOK, "login.html", nil)
		})
		// 登录操作
		userGroup.POST("/login", func(c *gin.Context) {
			//form表单数据:参数绑定
			var user User
			if err := c.ShouldBind(&user); err != nil { //出错
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 2,
					"msg":  err.Error(),
				})
				return
			}
			//成功绑定，然后查找用户是否存在
			result := db.Where("username = ?", user.Username).First(&user)
			if result.RowsAffected == 0 { //不存在
				c.JSON(http.StatusOK, gin.H{ //响应
					"code": 0,
					"msg":  "No such user",
				})
				return
			}
			//存在，登录成功
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Login successfully, welcome",
			})

			//重定向至主页

		})

	}

	r.Run()
}
