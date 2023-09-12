package main

import (
	"os"
	"rahul-yr/learn-go-simple-todo-crud/todo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// main web server
func main() {
	
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()

	r.Use(cors.Default())

	router := r.Group("/todo")
	todo.MountTodoRouter(router)
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
