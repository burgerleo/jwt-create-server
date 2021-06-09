package token

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	var ACCESS_SECRET = "jdnfksdmfksd"

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ACCESS_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TestCreatetoken(t *testing.T) {

	var jwt = Jwt{}
	err := jwt.GenerateToken(1)
	if err != nil {
		fmt.Println("err")
	}

	print(jwt.RetrieveToken())

	// assert.Equal(t, "cde",)
}
