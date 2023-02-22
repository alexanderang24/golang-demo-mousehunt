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

type trapController struct {
	service services.TrapService
}

func NewTrapController(service services.TrapService) *trapController {
	return &trapController{service}
}

func (c *trapController) GetAllTraps(ctx *gin.Context) {
	traps, err := c.service.GetAllTraps()
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

func (c *trapController) GetTrap(ctx *gin.Context) {
	var trap structs.Trap
	id, _ := strconv.Atoi(ctx.Param("id"))
	trap.ID = int64(id)

	trap, err := c.service.GetTrap(trap)
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

func (c *trapController) InsertTrap(ctx *gin.Context) {
	var trap structs.Trap

	err := ctx.ShouldBindJSON(&trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	trap, err = c.service.InsertTrap(trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success insert trap with ID: %d and name: %s", trap.ID, trap.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func (c *trapController) UpdateTrap(ctx *gin.Context) {
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
	trap, err = c.service.UpdateTrap(trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success update trap with ID: %d and name: %s", trap.ID, trap.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result})
}

func (c *trapController) DeleteTrap(ctx *gin.Context) {
	var trap structs.Trap
	id, _ := strconv.Atoi(ctx.Param("id"))
	trap.ID = int64(id)

	trap, err := c.service.DeleteTrap(trap)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("Success delete trap with ID: %d and name: %s", trap.ID, trap.Name)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}

func (c *trapController) BuyTrap(ctx *gin.Context) {
	// check trap exist
	var trap structs.Trap
	id, _ := strconv.Atoi(ctx.Param("id"))
	trap.ID = int64(id)
	trap, err := c.service.GetTrap(trap)
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
	var us services.UserService
	user, err := us.GetByUsername(username)

	// buy trap
	user, err = c.service.BuyTrap(trap, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := fmt.Sprintf("You bought %s that costs %d gold. You have %d gold now", trap.Name, trap.Price, user.Gold)
	ctx.JSON(http.StatusOK, gin.H{ "result": result })
}
