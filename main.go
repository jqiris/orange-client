package main

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

func main() {
	app := gin.Default()
	app.Use(cors.Default())
	app.Static("/", "web-mobile")
	app.Run(":18888")
}
