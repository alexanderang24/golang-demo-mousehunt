package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/structs"
	"net/http"
)

type huntHistoryController struct {
	service services.HuntHistoryService
}

func NewHuntHistoryController(service services.HuntHistoryService) *huntHistoryController {
	return &huntHistoryController{service}
}

func (c *huntHistoryController) GetAllHuntHistories(ctx *gin.Context) {
	histories, err := c.service.GetAllHistories()
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

func (c *huntHistoryController) DoHunt(ctx *gin.Context) {
	var history structs.HuntHistory

	err := ctx.ShouldBindJSON(&history)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	history, err = c.service.DoHunt(history)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert history")
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}