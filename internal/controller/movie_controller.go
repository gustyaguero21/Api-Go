package controller

import (
	"api-go/internal/models"
	"api-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


type MovieContoller struct{
	MS services.MovieService
}

func(mc *MovieContoller)AddDirectorController(ctx *gin.Context){
	var director models.Director


	if err:=ctx.ShouldBindJSON(&director);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error deserializing body":err.Error(),
		})
		return
	}

	created,createErr:=mc.MS.AddDirectorService(ctx,director)
	if createErr!=nil{
		ctx.JSON(http.StatusInternalServerError, createErr.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
    	"message": "Director agregado exitosamente",
    	"director": created,
	})
}

func(mc *MovieContoller)AddActorController(ctx *gin.Context){
	var actor models.Actor


	if err:=ctx.ShouldBindJSON(&actor);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error deserializing body":err.Error(),
		})
		return
	}

	created,createErr:=mc.MS.AddActorService(ctx,actor)
	if createErr!=nil{
		ctx.JSON(http.StatusInternalServerError, createErr.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
    	"message": "Actor agregado exitosamente",
    	"actor": created,
	})
}

