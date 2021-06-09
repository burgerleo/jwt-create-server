package service

import (
	"fmt"
	"jwt-generate-server/models"
	"jwt-generate-server/service/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiOutput struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

var apiSuccess = ApiOutput{
	Message: "Success",
	Status:  200,
}

func Home(ctx *gin.Context) {
	ctx.JSON(200, apiSuccess)
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
	fmt.Println(";;;;;;;;;;;;;;")
	var user models.User
	// 出現這個錯誤
	// [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 200
	// https://studygolang.com/articles/17745
	// ctx.BindJSON(&user)
	ctx.ShouldBind(&user)

	fmt.Println(user)

	var jwt = token.JWTToken
	err := jwt.GenerateToken(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			ApiOutput{
				Message: "Fail",
				Status:  http.StatusBadRequest,
				Data:    "Jwt generate error",
			})
		return
	}

	ctx.JSON(http.StatusOK,
		ApiOutput{
			Message: "Success",
			Status:  http.StatusOK,
			Data:    jwt,
		})
}

func JwtVerify(ctx *gin.Context) {
	ctx.JSON(200, apiSuccess)
}
