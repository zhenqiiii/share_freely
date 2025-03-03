package logic

import (
	"github.com/zhenqiiii/share_freely/gorm/mysql"
	"github.com/zhenqiiii/share_freely/models"
)

// 业务逻辑：功能模块

// 换一换功能
func FuncChange() (post *models.Post, err error) {
	post, err = mysql.GetRandomPost() //是否还会出现nil reference?
	if err != nil {
		return nil, err
	}
	return post, nil
}
