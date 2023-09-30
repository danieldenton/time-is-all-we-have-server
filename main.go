package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID      string `json: "id"`
	Name    string `json: "task"`
	Time    string `json: "time"`
	Minutes int    `json: "minutes"`
}

var tasks = []task{
	{ID: "1", Name: "meditaion", Time: "10:00 am", Minutes: 30},
	{ID: "2", Name: "bjj", Time: "11:45 am", Minutes: 120},
	{ID: "3", Name: "work", Time: "2:00 pm", Minutes: 240},
	{ID: "4", Name: "study", Time: "7:00 pm", Minutes: 90},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.Run("localhost:8080")
}
