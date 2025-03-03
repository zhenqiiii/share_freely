package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhenqiiii/share_freely/logic"
)

// 主页模块的控制台代码：Handler
func HomeHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 如果前后端暂时不分离，那么这个函数就得先渲染模板
		// 接着直接返回一些主页homepage需要的数据
		post, err := logic.HomePage()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "fetch post failed!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "success!",
			"data": post, // 建议用JSON格式再封装一下？咋封装忘了...
		})

	}

}
