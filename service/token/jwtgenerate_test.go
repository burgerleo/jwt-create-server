package token

import (
	"fmt"
	"jwt-generate-server/models"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dgrijalva/jwt-go"
)

var (
	jtoken = JwtToken{secret: "abcd", expireTime: 10}
	user     = models.User{UserId: 1, Name: "leo"}
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

func TestGenerateToken(t *testing.T) {
	err := jtoken.GenerateToken(user)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("----- JWT Token -----")
	fmt.Println(jtoken.RetrieveToken())
	assert.NoError(t, err)
}

func TestVerifyToken(t *testing.T) {
	TestGenerateToken(t)

	token, err := jtoken.VerifyToken(jtoken.RetrieveToken())
	if err != nil {
		fmt.Println(err)
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	fmt.Println(claims["user"])
	leo := claims["user"]

	assert.Equal(t, leo.(map[string]interface{})["name"], "leo")
	assert.True(t, token.Valid)
}
