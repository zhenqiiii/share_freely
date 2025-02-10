package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// token有效时长
const TokenExpireDuration = time.Hour * 24

// Secret: 用于加盐的字符串，密钥
var Secret = []byte("你拿什么抗争")

// MyClaims 自定义声明类型 内嵌jwt.RegisteredClaims
// 加入Username字段
type MyClaims struct {
	UserID               int64  `json:"userid"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准声明
}

// GenToken 生成JWT字符串
func GenToken(userid int64, username string) (string, error) {
	// 创建MyClaims声明
	claims := MyClaims{ // 复习一下结构体的定义方式
		userid,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //过期时间
			Issuer:    "share_freely",                                          // 发行者
		},
	}

	// 创建Token对象
	// func NewWithClaims(method SigningMethod, claims Claims, opts ...TokenOption) *Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 通过密钥获得token字符串
	// func (t *Token) SignedString(key interface{}) (string, error)
	return token.SignedString(Secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 这里我们是自定义的Claims结构体，需要用ParseWithClaims方法
	// 标准Claims的话可以直接用Parse
	var claims = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	// 效验token是否有效
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
