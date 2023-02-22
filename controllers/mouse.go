package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/structs"
	"net/http"
	"strconv"
)

func GetAllMice(ctx *gin.Context) {
	mice, err := services.GetAllMice()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"results": mice,
		})
	}
}

func GetMouse(ctx *gin.Context) {
	var mouse structs.Mouse
	id, _ := strconv.Atoi(ctx.Param("id"))
	mouse.ID = int64(id)

	mouse, err := services.GetMouse(mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": mouse,
		})
	}
}

func InsertMouse(ctx *gin.Context) {
	var mouse structs.Mouse

	err := ctx.ShouldBindJSON(&mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	mouse, err = services.InsertMouse(mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert mouse with ID: %d and name: %s", mouse.ID, mouse.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func UpdateMouse(ctx *gin.Context) {
	var mouse structs.Mouse

	err := ctx.ShouldBindJSON(&mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	mouse.ID = int64(id)
	mouse, err = services.UpdateMouse(mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success update mouse with ID: %d and name: %s", mouse.ID, mouse.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result})
}

func DeleteMouse(ctx *gin.Context) {
	var mouse structs.Mouse
	id, _ := strconv.Atoi(ctx.Param("id"))
	mouse.ID = int64(id)

	mouse, err := services.DeleteMouse(mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success delete mouse with ID: %d and name: %s", mouse.ID, mouse.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}