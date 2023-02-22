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

func GetAllLocations(ctx *gin.Context) {
	locations, err := services.GetAllLocations()
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

func GetLocation(ctx *gin.Context) {
	var location structs.Location
	id, _ := strconv.Atoi(ctx.Param("id"))
	location.ID = int64(id)

	location, err := services.GetLocation(location)
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

func InsertLocation(ctx *gin.Context) {
	var location structs.Location

	err := ctx.ShouldBindJSON(&location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	location, err = services.InsertLocation(location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert location with ID: %d and name: %s", location.ID, location.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func UpdateLocation(ctx *gin.Context) {
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
	location, err = services.UpdateLocation(location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success update location with ID: %d and name: %s", location.ID, location.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result})
}

func DeleteLocation(ctx *gin.Context) {
	var location structs.Location
	id, _ := strconv.Atoi(ctx.Param("id"))
	location.ID = int64(id)

	location, err := services.DeleteLocation(location)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success delete location with ID: %d and name: %s", location.ID, location.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func TravelToLocation(ctx *gin.Context) {
	// check location exist
	var location structs.Location
	id, _ := strconv.Atoi(ctx.Param("id"))
	location.ID = int64(id)

	location, err := services.GetLocation(location)
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

	// travel
	user, err = services.TravelToLocation(location, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("You travel to %s that costs %d gold. You have %d gold now", location.Name, location.TravelCost, user.Gold)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}