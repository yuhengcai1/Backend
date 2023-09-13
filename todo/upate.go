package todo

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func updateTodo(c1 *gin.Context) {
	var inputTodo Todo
	c1.BindJSON(&inputTodo)
	id := c1.Param("id")
	todoId, err := strconv.Atoi(id)
	needtobe.ID = todoId
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid"})
		return
	}
	if needtobe.Title == "" {
		c.JSON(400, gin.H{"error": "Needed title"})
		return
	}
	for i, todo := range todos {
		if todo.ID == todoId {
			todos[i] = needtobe
			c.JSON(200, needtobe)
			return
		}
	}
	c.JSON(404, gin.H{"error": "error for update"})
}
