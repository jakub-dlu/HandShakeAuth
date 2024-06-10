package repository

import (
	"HSAuth/models"
	"HSAuth/store"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func AddUser(newUserReq models.RegisterRequest, c *gin.Context) { //adds user to database
	var newUser models.User

	// Check if email already exists
	emailExists, err := DoesEmailExist(newUserReq.Email)
	if err != nil {
		fmt.Println("Failed to check if email exists:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email"})
		return
	}
	if emailExists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	publicKeyBytes, err := base64.StdEncoding.DecodeString(newUserReq.PublicKey) //decoding received string to bytes
	if err != nil {
		fmt.Println("Failed to decode public key:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid public key format"})
		return
	}

	newUser = models.User{
		ID:        uuid.New().String(), //unique ID
		Email:     newUserReq.Email,    //email from register request
		PublicKey: publicKeyBytes,      //public key decoded to bytes from register req
	}

	if err := store.DB.Create(&newUser).Error; err != nil { //create new user in database
		fmt.Println("Failed to add user:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Failed to add user:": err})
	} else {
		fmt.Println("User added successfully")
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
	}
}
