package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/dto"
	"net/http"
	"strconv"
)

func GetAllUsers(ctx *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"results": users,
		})
	}
}

func GetUser(ctx *gin.Context) {
	var user dto.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	user.ID = int64(id)

	user, err := services.GetUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"result": user,
		})
	}
}

func Register(ctx *gin.Context) {
	var user dto.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = services.Register(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Registration success! Welcome, Hunters %s!", user.Username)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func InsertUser(ctx *gin.Context) {
	var user dto.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = services.InsertUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert user with ID: %d and username: %s", user.ID, user.Username)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func UpdateUser(ctx *gin.Context) {
	var user dto.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	user.ID = int64(id)
	user, err = services.UpdateUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success update user with ID: %d and username: %s", user.ID, user.Username)
	ctx.JSON(http.StatusOK, gin.H{ "result": result})
}

func DeleteUser(ctx *gin.Context) {
	var user dto.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	user.ID = int64(id)

	user, err := services.DeleteUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success delete user with ID: %d and username: %s", user.ID, user.Username)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func Login(ctx *gin.Context) {
	var user dto.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := services.Login(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ "result": token })
}
