package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
)

func VerifyJWT(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, invalid token format",
		})
		return
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized, token verification failed",
			})
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, error parsing JWT",
		})
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, invalid token",
		})
		return
	}

	if token.Valid {
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, invalid token",
		})
		return
	}
}
