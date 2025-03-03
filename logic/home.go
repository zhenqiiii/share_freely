package logic

import (
	"github.com/zhenqiiii/share_freely/gorm/mysql"
	"github.com/zhenqiiii/share_freely/models"
)

// 业务逻辑函数：主页模块
func HomePage() (post *models.Post, err error) {
	// 从数据库中随机抽取一篇文章返回过去
	post, err = mysql.GetRandomPost()
	// 想想还有什么？
	// 考虑一下主页有啥功能
	// 想写一个天气，应该要从某个接口拉数据过来
	return post, err
}
