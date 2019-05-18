package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg": "Hello, world!",
	})
}