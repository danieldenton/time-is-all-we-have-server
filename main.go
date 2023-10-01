package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID        int    `json: "id"`
	Name      string `json: "task"`
	Time      string `json: "time"`
	Minutes   int    `json: "minutes"`
	Completed bool   `json: "completed"`
}

var tasks = []task{
	{ID: 1, Name: "meditation", Time: "10:00 am", Minutes: 30, Completed: false},
	{ID: 2, Name: "bjj", Time: "11:45 am", Minutes: 120, Completed: false},
	{ID: 3, Name: "work", Time: "2:00 pm", Minutes: 240, Completed: false},
	{ID: 4, Name: "study", Time: "7:00 pm", Minutes: 90, Completed: false},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func taskByName(c *gin.Context) {
	name := c.Param("name")
	task, err := getTaskByName(name)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func completeTask(c *gin.Context) {
	name, ok := c.GetQuery("name")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing name quesry parameter."})
		return
	}

	task, err := getTaskByName(name)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}

	if task.Completed == true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task already completed"})
	}

	task.Completed = true
	c.IndentedJSON(http.StatusOK, task)
}

func getTaskByName(name string) (*task, error) {
	for i, t := range tasks {
		if t.Name == name {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func createTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTask)
	router.GET("/tasks/:name", taskByName)
	router.PATCH("/complete", completeTask)
	router.Run("localhost:8080")
}
