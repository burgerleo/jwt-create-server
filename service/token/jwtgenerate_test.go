package token

import (
	"fmt"
	model "jwt-generate-server/models"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtToken = JwtToken{secret: "abcd", expireTime: 10}
	user     = model.User{UserId: 1, Name: "leo"}
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
	err := jwtToken.GenerateToken(user)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("----- JWT Token -----")
	fmt.Println(jwtToken.RetrieveToken())
	assert.NoError(t, err)
}

func TestVerifyToken(t *testing.T) {
	TestGenerateToken(t)

	token, err := jwtToken.VerifyToken(jwtToken.RetrieveToken())
	if err != nil {
		fmt.Println(err)
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	assert.Equal(t, claims["user_id"], float64(1))
	assert.True(t, token.Valid)
}
