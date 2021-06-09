package service

import (
	"fmt"
	"jwt-generate-server/service/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiOutput struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func Home(ctx *gin.Context) {
	ctx.JSON(200, ApiOutput{
		Message: "Success",
		Status:  200,
		Data:    "",
	})
}

func HellowLeo(ctx *gin.Context) {
	type Leo struct {
		UserId int    `json:"user_id"`
		Name   string `json:"name"`
	}
	leo := Leo{1, "leo"}

	ctx.JSON(http.StatusOK,
		ApiOutput{
			Message: "Success",
			Status:  http.StatusOK,
			Data:    leo,
		})
}

func GetJsonData(ctx *gin.Context) {
	type Leo struct {
		UserId int    `json:"user_id"`
		Name   string `json:"name"`
	}
	// json := make(map[string]interface{})
	json := Leo{}

	ctx.BindJSON(&json)

	ctx.JSON(http.StatusOK,
		ApiOutput{
			Message: "Success",
			Status:  http.StatusOK,
			Data:    json,
		})
}

func JwtGenerate(ctx *gin.Context) {
	json := make(map[string]interface{})

	ctx.BindJSON(&json)

	var jwt = token.Jwt{}
	err := jwt.GenerateToken(1)
	if err != nil {
		fmt.Println("err")
	}

	// print(jwt.RetrieveToken())

	ctx.JSON(http.StatusOK,
		ApiOutput{
			Message: "Success",
			Status:  http.StatusOK,
			Data:    jwt,
		})
}
