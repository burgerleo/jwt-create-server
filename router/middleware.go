package router

import (
	"jwt-generate-server/service/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// https://ithelp.ithome.com.tw/articles/10243831

func validateJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := token.JWTToken
		jwt := getTokenString(ctx.GetHeader("Authorization"))

		if _, err := token.VerifyToken(jwt); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "token expired or token validate fail",
				"status":  http.StatusUnauthorized,
			})
			ctx.Abort()
		}
	}
}

func getTokenString(s string) string {
	// s = Bearer eyJhbGci....
	// 取得第二個 string
	strArr := strings.Split(s, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
