package token

import (
	"jwt-generate-server/models"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtToken *JwtToken
	once     = &sync.Once{}
)

type Token interface {
	GenerateToken(models.User) error
	RetrieveToken() string
	VerifyToken(string) (*jwt.Token, error)
}

func InitToken(secret string, expireTime int) error {
	jwtToken := GetJwtToken()
	jwtToken.Init(secret, expireTime)
	return nil
}

// 取得 jwtToken
func GetJwtToken() *JwtToken {
	if jwtToken == nil {
		once.Do(func() {
			jwtToken = &JwtToken{}
		})
	}
	return jwtToken
}
