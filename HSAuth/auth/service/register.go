package service

import (
	"HSAuth/models"
	"HSAuth/store/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) { //binding json and send to func AddUser
	var newUser models.RegisterRequest //email + public key

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	repository.AddUser(newUser, c)
}
