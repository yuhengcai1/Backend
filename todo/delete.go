package todo

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	todoId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	for i, todo := range todos {
		if todo.ID == todoId {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(200, gin.H{"message": "todo deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"error": "todo not found"})
}
