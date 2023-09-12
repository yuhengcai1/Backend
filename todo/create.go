package todo

import "github.com/gin-gonic/gin"

var idocount int = 1

func createTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	if todo.Title == "" {
		c.JSON(400, gin.H{"error": "title is required"})
		return
	} else {
		// set the id
		todo.ID = idocount
		todos = append(todos, todo)
		idocount++
		c.JSON(200, todo)
	}
}
