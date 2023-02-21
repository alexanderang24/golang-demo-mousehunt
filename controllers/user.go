package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/structs"
	"net/http"
	"strconv"
)

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *userController {
	return &userController{service}
}

func (c *userController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
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

func (c *userController) GetUser(ctx *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	user.ID = int64(id)

	user, err := c.service.GetUser(user)
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

func (c *userController) Register(ctx *gin.Context) {
	var user structs.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = c.service.Register(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Registration success! Welcome, Hunters %s!", user.Username)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func (c *userController) InsertUser(ctx *gin.Context) {
	var user structs.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = c.service.InsertUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert user with ID: %d and username: %s", user.ID, user.Username)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var user structs.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	user.ID = int64(id)
	user, err = c.service.UpdateUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success update user with ID: %d and username: %s", user.ID, user.Username)
	ctx.JSON(http.StatusOK, gin.H{ "result": result})
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	user.ID = int64(id)

	user, err := c.service.DeleteUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success delete user with ID: %d and username: %s", user.ID, user.Username)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func (c *userController) Login(ctx *gin.Context) {
	var user structs.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = c.service.Login(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ "result": user.Token })
}
