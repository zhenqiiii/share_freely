package models

import "time"

// 用户模型
type User struct {
	UID      int64  `gorm:"primaryKey"`
	Username string `json:"username" form:"username" gorm:"type:varchar(100)"`
	Password string `json:"password" form:"password" gorm:"type:varchar(100)"`
}

// 注册时的参数
type ParamRegister struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// 登录时的参数
type ParamLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// 文章模型
type Post struct {
	PID        int64  `gorm:"primaryKey;autoIncrement"`
	Title      string `gorm:"type:varchar(255);not null"`
	Content    string `gorm:"type:text;not null"`
	IsOriginal bool   `gorm:"type:boolean"`
	Category   string `gorm:"type:varchar(100)"`
	UploaderID int64
	UploadTime time.Time `gorm:"autoCreateTime"`
}

// 上传时的参数:主要与文章有关，接收到后确认参数无误再创建文章实例，加入库中
// 用户信息由jwt确定
type ParamUpload struct {
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	Category   string `json:"category" form:"category"`
	IsOriginal bool   `json:"isoriginal" form:"isoriginal"`
}
