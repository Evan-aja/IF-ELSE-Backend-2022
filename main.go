package main

import (
	"ifelse/Controller"
	"ifelse/Database"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	//Database
	db := Database.Open()
	if db != nil {
		println("Nice, DB Connected")
	}

	// Gin Framework
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.SetTrustedProxies(
		[]string{
			os.Getenv("PROXY_1"),
			os.Getenv("PROXY_2"),
			os.Getenv("PROXY_3"),
		},
	)

	//CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Controll-Allow-Methods", "POST,OPTIONS,GET,PUT,DELETE,PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	//Routers
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to IF-ELSE 2022",
			"success": true,
		})
	})
	r.Group("/api")
	Controller.Register(db, r)
	Controller.User(db, r)

	//Server init
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
		return
	}
}
