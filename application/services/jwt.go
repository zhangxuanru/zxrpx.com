/*
@Time : 2020/7/13 14:59
@Author : zxr
@File : jwt
@Software: GoLand
*/
package services

import (
	"errors"
	"pix/application/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet       = errors.New("Token not active yet")
	TokenMalformed         = errors.New("That's not even a token")
	TokenInvalid           = errors.New("Couldn't handle this token:")
	SignKey                = "zxrStrive"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	AccountAuth
	jwt.StandardClaims
}

type AccountAuth struct {
	UserId   int    `json:"user_id"`
	UserPxId int    `json:"user_px_id"`
	UserName string `json:"user_name"`
	//Password string `json:"password"`
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return SignKey
}

//生成用户登录token
func (j *JWT) GenUserToken(user models.User) (token string, err error) {
	claims := CustomClaims{
		AccountAuth: AccountAuth{
			UserId:   user.Id,
			UserPxId: user.PxUid,
			UserName: user.UserName,
			//Password: user.Passwd,
		},
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),        // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
		},
	}
	token, err = j.CreateToken(claims)
	return
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token有效期
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
