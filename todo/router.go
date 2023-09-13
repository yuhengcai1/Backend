package todo

import "github.com/gin-gonic/gin"

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo = []Todo{}

func MountTodoRouter(c *gin.RouterGroup) {
	c.POST("/create", createTodo)
	c.GET("/read/:id", readTodo)
	c.POST("/update/:id", updateTodo)
	c.DELETE("/delete/:id", deleteTodo)
}
