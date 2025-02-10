package gorm

import (
	"errors"

	"github.com/zhenqiiii/share_freely/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

// 注册模块Register
// 1.判断用户是否存在
func CheckUserExist(username string) (err error) {
	var user models.User
	if result := db.Where("name = ?", username).First(&user); result.RowsAffected != 0 {
		return errors.New("用户已存在")
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
		return nil, errors.New("用户不存在")
	}
	if user2.Password != user.Password {
		return nil, errors.New("密码错误")
	}
	return user2, nil
}
