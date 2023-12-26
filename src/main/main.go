package main

import (
	"fmt"
	"os"
	"todolist-go/src/models"
	"todolist-go/src/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	engine := gin.Default()

	port := os.Getenv("PORT")

	r := models.Server{Engine: engine}
	routers.UserRouter(r, "/users")
	routers.TodoRouter(r, "/todos")
	fmt.Println("server run on port ", port)
	r.Engine.Run(":" + port)
}
