package main

import (
	"HSAuth/api"
	"HSAuth/store"
)

func main() {
	store.InitDB()
	api.InitRouter()
}
