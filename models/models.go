package models

// 用户模型
type User struct {
	UID      int64  `gorm:"primaryKey"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
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
	PID        int64 `gorm:"primaryKey"`
	UploaderID int64
}

// 上传时的参数
type ParamUpload struct {
}
