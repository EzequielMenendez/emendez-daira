package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type history struct{
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
	Operation string `json: "operation"`
	Result int `json: "result"`
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

	router.GET("/history", getHistory)
	router.POST("/history", postHistory)

	router.Run("localhost:3000")
}