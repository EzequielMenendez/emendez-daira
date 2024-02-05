package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type history struct{
	Date string `json: "date"`
	Operation string `json: "operation"`
	Result float64 `json: "result"`
}

var historys = []history{}

func getHistory(c *gin.Context){
	c.IndentedJSON(http.StatusOK, historys)
}

func postHistory(c *gin.Context){
	var newHistory history
	
	c.BindJSON(&newHistory)

	historys = append(historys, newHistory)

	c.IndentedJSON(http.StatusCreated, historys)
}

func main(){
	router := gin.Default()
	router.Use(corsMiddleware())

	router.GET("/history", getHistory)
	router.POST("/history", postHistory)

	router.Run("localhost:3000")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}