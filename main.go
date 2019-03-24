package main

import (
	"cocotor.com/api"
	"cocotor.com/api/common"
	"cocotor.com/api/news"
	"cocotor.com/middleware/cors"
	"cocotor.com/middleware/jwt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {
    // Logging to a file.
    f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
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
	
	r.Run(":8080")
}