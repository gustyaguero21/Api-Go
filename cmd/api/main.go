package main

import (
	"api-go/cmd/config"
	"api-go/internal/router"
)

func main(){
	config.LoadEnv()
	
	router.StartServer()
}