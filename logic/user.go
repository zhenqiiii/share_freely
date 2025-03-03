package logic

import (
	"github.com/zhenqiiii/share_freely/gorm/mysql"
	"github.com/zhenqiiii/share_freely/models"
	"github.com/zhenqiiii/share_freely/pkg/jwt"
)

// 业务逻辑：用户模块
// 通过参数校验后注册
func Register(param *models.ParamRegister) (err error) {
	// 1.判断用户是否存在
	if err = mysql.CheckUserExist(param.Username); err != nil {
		return err
	}
	// 2.通过雪花算法生成UID
	// 3.构造完整User实例
	// 4.插入数据库
	return nil
}

func Login(param *models.ParamLogin) (token string, err error) {
	// 1.参数校验通过后构造User实例(参数校验属于控制台函数范畴)
	user := models.User{
		UID:      0,
		Username: param.Username,
		Password: param.Password,
	}
	// 2.查询数据库，未查询到则返回空token+错误，查询到则返回token+nil
	RealUser, err := mysql.Login(&user)
	if err != nil {
		return "", err
	}
	// 查询到且信息正确，生成Token
	return jwt.GenToken(RealUser.UID, RealUser.Username)

}
