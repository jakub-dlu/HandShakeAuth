package service

import (
	"HSAuth/models"
	"HSAuth/store"
	"HSAuth/store/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetServersPublicKey(c *gin.Context) { //sends server's public key to user
	var k models.KeysModel
	if err := store.DB.First(&k); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"publicKey": k.PublicKey})
}

func LoginWithSharedSecret(c *gin.Context) { //binding json and sending it to func verifyUser
	var loginDetails models.LoginRequest //email + shared secret

	err := c.ShouldBindJSON(&loginDetails)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	repository.VerifyUser(loginDetails, c)
}
