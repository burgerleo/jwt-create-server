package token

import (
	"fmt"
	"jwt-generate-server/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	Token      string `json:"token"`
	secret     string
	expireTime int
}

func (j *JwtToken) Init(secret string, expire int) {
	fmt.Println(secret)

	j.setSecret(secret)

	j.setExpireTime(expire)
}

func (j *JwtToken) setSecret(secret string) {
	j.secret = secret
}

func (j *JwtToken) setExpireTime(expire int) {
	j.expireTime = expire
}

// https://learn.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr/
func (j *JwtToken) GenerateToken(user models.User) error {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(j.expireTime)).Unix()
	atClaims["user"] = user

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(j.secret))
	if err != nil {
		return err
	}

	j.Token = token

	return nil
}

func (j *JwtToken) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (j *JwtToken) RetrieveToken() string {
	return j.Token
}
