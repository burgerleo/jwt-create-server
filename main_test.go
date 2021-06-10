package main

import (
	"fmt"
	"jwt-generate-server/models"
	"jwt-generate-server/service/token"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://mileslin.github.io/2020/06/Golang/Unit-testing-HTTP-servers-with-Gin/
func Test_setupRouter(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/leo", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
	assert.Contains(t, w.Body.String(), "leo")
}

func TestGetToken(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件

	// set JSON body
	jsonParam := `{"user_id":1,"name":"leo"}`

	// Mock HTTP Request and it's return
	req, _ := http.NewRequest("POST", "/jwt", strings.NewReader(string(jsonParam)))
	// 傳送 json 要修改 header type
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
	assert.Contains(t, w.Body.String(), "Success")
}

func TestVerifyToken(t *testing.T) {
	jwt := token.GetJwtToken()
	jwt.GenerateToken(models.User{UserId: 1, Name: "leo"})

	router := setupRouter()
	// 取得 ResponseRecorder 物件
	w := httptest.NewRecorder()

	// Mock HTTP Request and it's return
	req, _ := http.NewRequest("GET", "/jwt/verify", nil)
	req.Header.Set("Authorization", "Bearer "+jwt.RetrieveToken())
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Success")
}
