package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/share_freely/logic"
	"github.com/zhenqiiii/share_freely/models"
)

// 功能模块的控制台代码：Handler 说白了就是路由处理函数
// 这些处理函数都和普通处理函数格式一样（中间件）

// =========================主页功能===============================//

// 主页的换一换功能
// 无请求参数
// 返回随机文章
func ChangePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 调用业务逻辑函数
		post, err := logic.FuncChange()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "fail to fetch post",
				"error": err,
			})
			return
		}
		// 返回文章
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "fetch post successfully!",
			"data": post,
		})
	}
}

// =========================上传==================================//
// 上传文章功能
func UploadPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收参数
		var p models.ParamUpload
		err := c.ShouldBind(&p)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "something went wrong with the param",
				"error": err,
			})
			return
		}
		// 获取token
		uid, _ := c.Get("userid")
		// 调用业务逻辑
		err = logic.UploadPost(p, uid.(int64))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "something went wrong",
				"error": err,
			})
			return
		}
		// 上传成功
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "upload successfully",
		})

	}
}

// 查看我的上传内容
func MyUploads() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取uid
		uid, _ := c.Get("userid")
		// 执行业务逻辑函数
		posts, err := logic.ViewMyUpload(uid.(int64))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"msg":   "fail to fetch Uploads!",
				"error": err,
			})
			return
		}
		// 未出错，返回上传内容
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  strconv.Itoa(len(posts)) + "posts in total",
			"data": posts,
		})

	}
}

// =========================搜索==================================//
