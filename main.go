package main

import (
	"cocotor.com/api"
	"cocotor.com/api/common"
	"cocotor.com/api/news"
	"cocotor.com/api/torch"
	"cocotor.com/middleware/cors"
	"cocotor.com/middleware/jwt"
	"github.com/gin-gonic/gin"
	"log"
)


func main() {
	r := gin.Default()
	r.Use(cors.Cors())
	log.Println("路由开始....")
	r.POST("/login", common.Login)
	r.POST("/register", common.Register)

	app := r.Group("/app")
	app.Use(jwt.JWTAuth())
	{
		app.GET("/test", api.Test)
		app.GET("/news", news.News)
	}

	torchs := r.Group("/torch")
	torchs.Use(jwt.JWTAuth())
	{
		torchs.GET("/plants", torch.AllPlants)
		torchs.GET("/plant/info", torch.PlantInfo)
		torchs.GET("/plant/status", torch.PlantStatus)
	}
	
	r.Run(":8080")
}