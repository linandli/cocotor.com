package torch

import (
	"fmt"
	"reflect"
	"net/http"
	"github.com/gin-gonic/gin"
	"cocotor.com/utils/torchlib"
)

func Loads(c *gin.Context) {
	// req := requests.Requests()
	// resp,_ := req.Get("http://115.28.129.229:8080/powerMobileApp/resources/", requests.Auth{"asmcos","password...."})
	  // fmt.Println(resp.Text())
	plantId := c.DefaultQuery("plantid", "1")

	fmt.Println("plantId===========>", plantId)
	r := torchlib.Get("org/plants", nil, 60)
	fmt.Println("r------------>", reflect.TypeOf(r).String())
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg": r,
	})
}