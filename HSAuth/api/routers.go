package api

import (
	"HSAuth/auth/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() { //initialisation of router gin.Default
	r := gin.Default()
	r.POST("/register", service.Register)
	r.GET("/getPublicKey", service.GetServersPublicKey)
	r.POST("/login", service.LoginWithSharedSecret)
	r.Run()
}
