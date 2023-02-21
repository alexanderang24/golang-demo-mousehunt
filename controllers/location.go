package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/structs"
	"net/http"
	"strconv"
)

type locationController struct {
	service services.LocationService
}

func NewLocationController(service services.LocationService) *locationController {
	return &locationController{service}
}

func (c *locationController) GetAllLocations(ctx *gin.Context) {
	locations, err := c.service.GetAllLocations()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"results": locations,
		})
	}
}

func (c *locationController) GetLocation(ctx *gin.Context) {
	var location structs.Location
	id, _ := strconv.Atoi(ctx.Param("id"))
	location.ID = int64(id)

	location, err := c.service.GetLocation(location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": location,
		})
	}
}

func (c *locationController) InsertLocation(ctx *gin.Context) {
	var location structs.Location

	err := ctx.ShouldBindJSON(&location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	location, err = c.service.InsertLocation(location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert location with ID: %d and name: %s", location.ID, location.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func (c *locationController) UpdateLocation(ctx *gin.Context) {
	var location structs.Location

	err := ctx.ShouldBindJSON(&location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	location.ID = int64(id)
	location, err = c.service.UpdateLocation(location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success update location with ID: %d and name: %s", location.ID, location.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result})
}

func (c *locationController) DeleteLocation(ctx *gin.Context) {
	var location structs.Location
	id, _ := strconv.Atoi(ctx.Param("id"))
	location.ID = int64(id)

	location, err := c.service.DeleteLocation(location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success delete location with ID: %d and name: %s", location.ID, location.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}