package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/structs"
	"net/http"
	"strconv"
)

type mouseController struct {
	service services.MouseService
}

func NewMouseController(service services.MouseService) *mouseController {
	return &mouseController{service}
}

func (c *mouseController) GetAllMice(ctx *gin.Context) {
	mice, err := c.service.GetAllMice()
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

func (c *mouseController) GetMouse(ctx *gin.Context) {
	var mouse structs.Mouse
	id, _ := strconv.Atoi(ctx.Param("id"))
	mouse.ID = int64(id)

	mouse, err := c.service.GetMouse(mouse)
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

func (c *mouseController) InsertMouse(ctx *gin.Context) {
	var mouse structs.Mouse

	err := ctx.ShouldBindJSON(&mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	mouse, err = c.service.InsertMouse(mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert mouse with ID: %d and name: %s", mouse.ID, mouse.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func (c *mouseController) UpdateMouse(ctx *gin.Context) {
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
	mouse, err = c.service.UpdateMouse(mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success update mouse with ID: %d and name: %s", mouse.ID, mouse.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result})
}

func (c *mouseController) DeleteMouse(ctx *gin.Context) {
	var mouse structs.Mouse
	id, _ := strconv.Atoi(ctx.Param("id"))
	mouse.ID = int64(id)

	mouse, err := c.service.DeleteMouse(mouse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success delete mouse with ID: %d and name: %s", mouse.ID, mouse.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

//func GetBooksByCategory(c *gin.Context) {
//	var result gin.H
//
//	var cat structs.Mouse
//	id, _ := strconv.Atoi(c.Param("id"))
//	cat.ID = int64(id)
//
//	cats, err := repository.GetBooksByCategory(database.DbConnection, cat)
//	if err != nil {
//		result = gin.H{ "error": err }
//	} else {
//		result = gin.H{ "result": cats }
//	}
//
//	c.JSON(http.StatusOK, result)
//}