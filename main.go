package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	Id string `json:"id"`
	Task string `json:"task"`
	Status string `json:"status"`
}

var todos = []todo {
	{Id:"1", Task: "Cleaning room", Status: "Not started"},
	{Id:"2", Task: "Writing email", Status: "On going"},
	{Id:"3", Task: "Feeding a cat", Status: "Not started"},
	{Id:"4", Task: "Learning Go lang", Status: "Not started"},
}

func getAllTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTasks(context *gin.Context) {
	var newTask todo

	if err := context.BindJSON(&newTask); err != nil {
		return 
	}

	todos = append(todos, newTask)
	context.IndentedJSON(http.StatusCreated, newTask)
}

func detailTask(context *gin.Context) {
	id := context.Param("id")
	for _, a := range todos {
		if a.Id == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No task found."})
}

func main() {
	fmt.Println("Running a server in localhost:9000")
	router := gin.Default()
	router.GET("/tasks", getAllTasks)
	router.POST("/tasks", addTasks)
	router.GET("/tasks/:id", detailTask)
	router.Run("localhost:9000")
}