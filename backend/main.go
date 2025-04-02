package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "RadheShyam, SitaRam, Joi ShivShakti: Welcome to Golang Gin backend!"})
	})

	fmt.Println("Server running on port 8080")
	log.Fatal(r.Run(":8080"))
}
