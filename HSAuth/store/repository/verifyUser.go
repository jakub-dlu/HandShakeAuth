package repository

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"

	"HSAuth/auth/keys"
	"HSAuth/models"
	"HSAuth/store"
)

func VerifyUser(loginDetails models.LoginRequest, c *gin.Context) { //verify login data
	var u models.User
	if err := store.DB.First(&u, "Email = ?", loginDetails.Email).Error; err != nil { //looking for user with received email
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	sharedSecretBytes, err := base64.StdEncoding.DecodeString(loginDetails.SharedSecret) //decoding received shared secret to bytes
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid shared secret format"})
		return
	}

	var k models.KeysModel
	if err := store.DB.First(&k).Error; err != nil { //get server`s key pair
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve server keys"})
		return
	}

	areSharedSecretsMatch, err := keys.CompareSharedSecrets(sharedSecretBytes, u.PublicKey, k) //compare both of shared secrets -> bool
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error comparing shared secrets"})
		return
	}

	if areSharedSecretsMatch { //check if shared secrets are the same
		c.IndentedJSON(http.StatusOK, gin.H{"message": "success!"})

		//================================
		//here you can add your own logic, for example access token or jwt
		//================================

		return
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
}
