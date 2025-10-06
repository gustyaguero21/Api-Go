package router

import (
	"api-go/cmd/config"

	"github.com/gin-gonic/gin"
)

func StartServer(){
	router:=gin.Default()

	URLMapping(router)

	router.Run(config.Port)
}