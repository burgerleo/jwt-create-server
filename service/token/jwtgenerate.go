package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	Token string `json:"token"`
}

func (j *Jwt) GenerateToken(userid uint64) error {
	var ACCESS_SECRET = "jdnfksdmfksd"
	var err error
	//Creating Access Token

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(ACCESS_SECRET))
	if err != nil {
		return err
	}

	j.Token = token

	return nil
}

func (j *Jwt) RetrieveToken() string {
	return j.Token
}
