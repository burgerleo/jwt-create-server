package token

import (
	"jwt-generate-server/models"

	"github.com/dgrijalva/jwt-go"
)

var JWTToken JwtToken

type Token interface {
	GenerateToken(models.User) error
	RetrieveToken() string
	VerifyToken(string) (*jwt.Token, error)
}

func InitToken(secret string, expireTime int) error {
	JWTToken.Init(secret, expireTime)
	return nil
}
