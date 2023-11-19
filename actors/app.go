package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var queue JobQueue = make(chan int)
var worker_pool = WorkerPool{
	concurrency: 3,
	Queue:       queue,
}

func CreateApp() *gin.Engine {
	r := gin.Default()
	r.GET("/health", health)
	r.POST("/greet", greet)
	return r
}

func health(ctx *gin.Context) {
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{"alive": "true"},
	)
}

func greet(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"msg": fmt.Sprintf("Hi %s!", user.Name),
		},
	)
}

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
