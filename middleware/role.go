package middleware

import (
	"github.com/gin-gonic/gin"
	"golang-demo-mousehunt/util"
	"net/http"
)

func AdminOnly(ctx *gin.Context) {
	_, jwtRole, err := util.ExtractClaims(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if jwtRole != "admin" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "forbidden access to API",
		})
		return
	}
	return
}
