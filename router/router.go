package router

import (
	"jwt-generate-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	authrized := r.Group("/")

	apiRoot := authrized.Group("/")
	{
		apiRoot.GET("home", service.Home)
		apiRoot.GET("leo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "leo",
				"status":  200,
			})
		})
		apiRoot.GET("leo2", service.HellowLeo)
		apiRoot.POST("json", service.GetJsonData)
	}

	tokenRouter := authrized.Group("/jwt")
	{
		tokenRouter.POST("", service.JwtGenerate)
		verifyRouter := tokenRouter.Group("/verify")
		{
			verifyRouter.Use(validateJwt())
			verifyRouter.GET("", service.JwtVerify)
		}
	}
}
