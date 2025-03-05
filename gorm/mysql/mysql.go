package mysql

import (
	"errors"

	"github.com/zhenqiiii/share_freely/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// =================Connection========================//
var db *gorm.DB

// 连接数据库并自动迁移模型
func ConnectToDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/share_freely?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}
	db.AutoMigrate(&models.User{}, &models.Post{})

	return nil
}

// // 检查数据库连接（防御性编程）
// func CheckConnection() error {
// 	if db == nil {
// 		return errors.New("database is not initialized")
// 	}
// }

// =======================User=========================//

// 注册模块Register
// 1.判断用户是否存在
func CheckUserExist(username string) (err error) {
	var user models.User
	if result := db.Where("username = ?", username).First(&user); result.RowsAffected != 0 {
		return errors.New("user already exists")
	}
	return nil
}

// 2.注册成功后插入数据库
func InsertUser(user *models.User) (err error) {
	if result := db.Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

// 登录模块Login
// 这里直接写一个服务于登录功能的数据库操作函数，查询到则返回user实例
// 否则返回err，在err中提示信息
func Login(user *models.User) (user2 *models.User, err error) {
	result := db.Where("username = ?", user.Username).First(&user2)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not exists")
	}
	if user2.Password != user.Password {
		return nil, errors.New("wrong pwd")
	}
	return user2, nil
}

// ======================function===============================//

// GetRandomPost:返回随机抽取的文章
func GetRandomPost() (RandomPost *models.Post, err error) {
	// 从数据库中随机查询文章
	result := db.Take(RandomPost)
	if result.Error != nil {
		return nil, result.Error
	}
	return RandomPost, nil
}

// InsertPost:向数据库加入文章
func InsertPost(post models.Post) error {
	// nil reference:可能由于db没有初始化成功
	result := db.Create(&post)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// GetMyUpload:获取我的上传
func GetMyUpload(uid int64) (posts []models.Post, err error) {
	// 根据uid查询,查找到的文章存入posts切片中
	result := db.Where("uploader_id = ?", uid).Find(&posts)
	if err = result.Error; err != nil {
		return nil, err
	}
	return posts, nil
}
