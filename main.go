package main

import (
	"fmt"

	"github.com/zhenqiiii/share_freely/gorm/mysql"
	"github.com/zhenqiiii/share_freely/router"
)

func main() {
	// 连接数据库
	err := mysql.ConnectToDB()
	if err != nil {
		fmt.Printf("连接数据库失败，错误：%v", err)
	}
	// 初始化路由
	r := router.SetupRouter("")

	// 小马快跑
	err = r.Run(":8080")
	if err != nil {
		fmt.Printf("连接服务器失败，错误：%v", err)
	}

}
