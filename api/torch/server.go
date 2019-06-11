package torch

import (
	"log"
	// "reflect"
	"net/http"
	"github.com/gin-gonic/gin"
	"cocotor.com/utils/torchlib"
)

func AllPlants(c *gin.Context) {
	r := torchlib.Get("org/plants/info/more", nil, 60)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg": r,
	})
}

func PlantInfo(c *gin.Context) {
	plant := c.DefaultQuery("plantid", "1")
	log.Println("plant-------->", plant)
	var plantId = map[string]string{"plantid": plant}
	r := torchlib.Get("org/plants/" + plant+ "/units" , plantId, 60)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg": r,
	})
}

func PlantStatus(c *gin.Context) {
	r := torchlib.Get("distribute/status", nil, 60)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg": r,
	})
}



