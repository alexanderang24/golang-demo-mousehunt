package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/dto"
	"golang-demo-mousehunt/dto/response"
	"golang-demo-mousehunt/services"
	"golang-demo-mousehunt/util"
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
	result := fmt.Sprintf("Registration success! Welcome, Hunter %s!", user.Username)
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

func GetMyInfo(ctx *gin.Context) {
	user := util.GetUserFromJWT(ctx)

	location, err := services.GetLocation(dto.Location{ID: user.LocationID})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "error when getting location data",
		})
		return
	}

	trap, err := services.GetTrap(dto.Trap{ID: user.TrapID})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "error when getting trap data",
		})
		return
	}

	var info = response.MyInfo{
		Username:     user.Username,
		Password:     user.Password,
		LocationName: location.Name,
		Gold:         user.Gold,
		Trap:         response.MyTrap{
			Name:     trap.Name,
			MinPower: trap.MinPower,
			MaxPower: trap.MaxPower,
		},
	}

	ctx.JSON(http.StatusOK, gin.H{ "result": info })
}

func GetCurrentLocationInfo(ctx *gin.Context) {
	user := util.GetUserFromJWT(ctx)

	location, err := services.GetLocation(dto.Location{ID: user.LocationID})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "error when getting location data",
		})
		return
	}

	mice, err := services.GetAllMiceInLocation(location.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "error when getting mice data",
		})
		return
	}

	var miceInHere []response.MiceInHere
	for _, mouse := range mice {
		miceInHere = append(miceInHere, response.MiceInHere{
			Name:        mouse.Name,
			Description: mouse.Description,
			MinPower:    mouse.MinPower,
			MaxPower:    mouse.MaxPower,
			GoldReward:  mouse.Gold,
		})
	}

	var info = response.LocationInfo{
		LocationName:        location.Name,
		LocationDescription: location.Description,
		TravelCost:          location.TravelCost,
		Mice:                miceInHere,
	}

	ctx.JSON(http.StatusOK, gin.H{ "result": info })
}