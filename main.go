package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	ID       int    `json:"id" form:"id" gorm:"autoIncrement"` //注册时由数据库指定：自增
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 文章模型
type Post struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	// Uploader   User      `gorm:"foreignKey:UploaderID"`
	UploaderID int       `json:"uploaderid"`
	UploadTime time.Time `json:"uploadtime" gorm:"autoCreateTime"`
	Title      string    `json:"title" form:"title" gorm:"type:varchar(255);not null" binding:"required"`
	Content    string    `json:"content" form:"content" gorm:"type:text;not null" binding:"required"`
	IsOriginal bool      `json:"isoriginal" form:"isoriginal" gorm:"type:boolean" binding:"required"`
}

func main() {
	//mysql database
	dsn := "root:root@tcp(127.0.0.1:3306)/share_freely?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//模型迁移
	db.AutoMigrate(&User{}, &Post{})

	var currentUser User //记录当前登录的用户

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
		//忘写密码错误场景了
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
			currentUser = user
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Login successfully, welcome",
			})

			//重定向至主页

		})

	}

	//home:主页模块
	homeGroup := r.Group("/home")
	{
		//访问主页
		homeGroup.GET("/main", func(c *gin.Context) {
			//渲染模版

			//获得随机文章
			var post Post
			result := db.Take(&post)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 0,
					"msg":  result.Error.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Successfully acquired",
				"data": post,
			})
		})

		//点击换一篇按钮
		homeGroup.GET("/change", func(c *gin.Context) {
			//获得随机文章
			var post Post
			result := db.Take(&post)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 0,
					"msg":  result.Error.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Successfully acquired",
				"data": post,
			})
		})

	}

	//功能模块-上传，查看我的上传
	funcGroup := r.Group("/func")
	{
		//点击上传按钮
		// 使用querystring发送username
		funcGroup.GET("/upload", func(c *gin.Context) {
			// 跳转至上传界面
			// c.HTML(http.StatusOK, "upload.html", nil)
			// username := c.Query("username")
			c.JSON(http.StatusOK, gin.H{ //响应
				"code": 1,
				"msg":  "username:" + currentUser.Username,
			})
		})
		// 确认上传
		funcGroup.POST("/upload", func(c *gin.Context) {
			// form表单提交内容：Title Content IsOriginal
			// shouldbind,利用binding:"required"机制判断前端填写内容是否完整
			var post Post
			if err := c.ShouldBind(&post); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 2,
					"msg":  err.Error(),
				})
				return
			}
			//上传内容完整，补全上传者信息
			post.UploaderID = currentUser.ID

			//写入数据库
			result := db.Create(&post)
			if result.Error != nil { //Fail
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 0,
					"msg":  result.Error.Error(),
				})
				return
			}
			//写入成功
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Successfully upload",
			})
		})
		//查看我的上传
		funcGroup.GET("/myuploads", func(c *gin.Context) {
			var posts []Post
			result := db.Where("uploader_id = ?", currentUser.ID).Find(&posts)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 0,
					"msg":  "Failed to search",
					"data": nil,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  strconv.Itoa(int(result.RowsAffected)) + "uploads in total",
				"data": posts,
			})
		})
	}

	r.Run()
}
