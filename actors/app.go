package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MAX_CONCURRENCY = 4
)

var queue JobQueue = make(chan int)
var worker_pool = WorkerPool{
	concurrency: MAX_CONCURRENCY,
	Queue:       queue,
	// set ouputs queue to nil, so we don't have to dequeue the messages
}

func main_web() {
	r := CreateApp()
	r.Run()
}

func CreateApp() *gin.Engine {
	go func() { worker_pool.Run() }()

	r := gin.Default()
	r.GET("/health", health)
	r.POST("/user", create_user)
	return r
}

func health(ctx *gin.Context) {
	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{"alive": "true"},
	)
}

func create_user(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	worker_pool.Queue <- user.Id

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
