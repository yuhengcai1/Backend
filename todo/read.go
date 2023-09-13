package todo

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func readTodo(c *gin.Context) {
	id := c.Param("id")
	todoId, _:= strconv.Atoi(id)
	for _, todo := range todos {
		// if the id matches the id in the todo
		if todo.ID == todoId {
			c.JSON(200, todo)
			return
		}
	}
	c.JSON(404, gin.H{"error": "None"})
}

func listTodos(c *gin.Context) {
	c.JSON(200, todos)
}
