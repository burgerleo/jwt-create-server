package token

import (
	model "jwt-generate-server/models"

	"github.com/dgrijalva/jwt-go"
)

var JWTToken JwtToken

type Token interface {
	GenerateToken(model.User) error
	RetrieveToken() string
	VerifyToken(string) (*jwt.Token, error)
}

func InitToken(secret string, expireTime int) error {
	JWTToken.Init(secret, expireTime)
	return nil
}
