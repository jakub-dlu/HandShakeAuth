package store

import (
	"HSAuth/auth/keys"
	"HSAuth/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() { // Initialize users database
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{}) //open database
	if err != nil {
		log.Fatalf("failed to connect to users database: %v", err)
	}

	fmt.Println("Users database connected successfully")

	//migrate user model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate users database: %v", err)
	}

	fmt.Println("Users database migrated successfully")

	//migrate key pair model
	if err := db.AutoMigrate(&models.KeysModel{}); err != nil {
		log.Fatalf("failed to migrate users database: %v", err)
	}

	//if key pair doesn't exist we need to create it
	var k models.KeysModel
	if err := db.First(&k).Error; err != nil {
		k.PublicKey, k.PrivateKey = keys.GenerateKeyPair()

		if err := db.Create(&k).Error; err != nil {
			log.Fatalf("failed to create keys: %v", err)
		}

		fmt.Println("Succesfully generated new key pair for server.")
	}

	fmt.Println("KeysModel migrated successfully")
	DB = db
}
