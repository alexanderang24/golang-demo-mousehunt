package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/util"
	"net/http"
)

func GetAllHuntHistories(ctx *gin.Context) {
	// get user
	user := util.GetUserFromJWT(ctx)

	// get histories for that user
	histories, err := services.GetAllHistories(user)
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
	// get user
	user := util.GetUserFromJWT(ctx)

	// do hunt
	response, err := services.DoHunt(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// set response message
	if response.Success {
		result := fmt.Sprintf("Your hunt in %s was successful! You encountered %s with %s power, luckily your trap generated %s power. You gained %s gold in your hunt, now you have %s gold",
			response.Location, response.MouseName, response.MousePower, response.TrapPower, response.GoldGained, response.GoldTotal)
		ctx.JSON(http.StatusOK, gin.H{ "result": result })
	} else {
		result := fmt.Sprintf("Unfortunately, your hunt in %s was failed! You encountered %s with %s power, but your trap only generated %s power.",
			response.Location, response.MouseName, response.MousePower, response.TrapPower)
		ctx.JSON(http.StatusOK, gin.H{ "result": result })
	}
}