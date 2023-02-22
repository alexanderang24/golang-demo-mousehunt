package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/middleware"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/structs"
	"net/http"
	"strconv"
)

func GetAllTraps(ctx *gin.Context) {
	traps, err := services.GetAllTraps()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"results": traps,
		})
	}
}

func GetTrap(ctx *gin.Context) {
	var trap structs.Trap
	id, _ := strconv.Atoi(ctx.Param("id"))
	trap.ID = int64(id)

	trap, err := services.GetTrap(trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": trap,
		})
	}
}

func InsertTrap(ctx *gin.Context) {
	var trap structs.Trap

	err := ctx.ShouldBindJSON(&trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	trap, err = services.InsertTrap(trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert trap with ID: %d and name: %s", trap.ID, trap.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func UpdateTrap(ctx *gin.Context) {
	var trap structs.Trap

	err := ctx.ShouldBindJSON(&trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	trap.ID = int64(id)
	trap, err = services.UpdateTrap(trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success update trap with ID: %d and name: %s", trap.ID, trap.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result})
}

func DeleteTrap(ctx *gin.Context) {
	var trap structs.Trap
	id, _ := strconv.Atoi(ctx.Param("id"))
	trap.ID = int64(id)

	trap, err := services.DeleteTrap(trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success delete trap with ID: %d and name: %s", trap.ID, trap.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func BuyTrap(ctx *gin.Context) {
	// check trap exist
	var trap structs.Trap
	id, _ := strconv.Atoi(ctx.Param("id"))
	trap.ID = int64(id)
	trap, err := services.GetTrap(trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get user
	username, _, err := middleware.ExtractClaims(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := services.GetByUsername(username)

	// buy trap
	user, err = services.BuyTrap(trap, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("You bought %s that costs %d gold. You have %d gold now", trap.Name, trap.Price, user.Gold)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}
