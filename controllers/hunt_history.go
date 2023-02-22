package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/structs"
	"net/http"
)

func GetAllHuntHistories(ctx *gin.Context) {
	histories, err := services.GetAllHistories()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"results": histories,
		})
	}
}

func DoHunt(ctx *gin.Context) {
	var history structs.HuntHistory

	err := ctx.ShouldBindJSON(&history)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	history, err = services.DoHunt(history)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert history")
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}