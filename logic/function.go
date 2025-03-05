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

// 上传功能
// 参数：1.上传参数 2.从token中获取的UploaderID
func UploadPost(param models.ParamUpload, uid int64) error {
	// 创建post实例
	post := models.Post{
		Title:      param.Title,
		Content:    param.Content,
		IsOriginal: param.IsOriginal,
		Category:   param.Category,
		UploaderID: uid,
	}
	// 插入数据库
	err := mysql.InsertPost(post)
	if err != nil {
		return err
	}

	return nil
}

// 查看我的上传
func ViewMyUpload(uid int64) (posts []models.Post, err error) {
	// 调用dao层函数查询
	posts, err = mysql.GetMyUpload(uid)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return posts, nil
}
