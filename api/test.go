package api

import (
	"net/http"
	// "time"

	// jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	// cocotorjwt "cocotor.com/middleware/jwt"
	// "cocotor.com/model"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg": "Hello, world!",
	})
}