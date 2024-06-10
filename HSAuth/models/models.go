package models

import "gorm.io/gorm"

type RegisterRequest struct {
	Email     string `json:"email"`
	PublicKey string `json:"publickey"` //client has to generate it
}

type LoginRequest struct {
	Email        string `json:"email"`
	SharedSecret string `json:"sharedSecret"` //client has to calculate it by his private key and server's pub key
}

type User struct {
	gorm.Model
	ID        string //generating automatically by uuid
	Email     string
	PublicKey []byte //get it from register request
}

type KeysModel struct { //model of key pair for SERVER
	PrivateKey []byte
	PublicKey  []byte
}
