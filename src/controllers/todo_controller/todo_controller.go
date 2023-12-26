package todo_controller

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	GetAllTodoWithUserId(*gin.Context)
}

type todoController struct {
	db *sql.DB
}

type Todolist struct {
	Id          int    `json:"id"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Title       string `json:"title"`
	UserId      int    `json:"user_id"`
}

func NewTodoController(db *sql.DB) TodoController {
	return todoController{db: db}
}

func (t todoController) GetAllTodoWithUserId(c *gin.Context) {
	userId := c.Query("user_id")
	fmt.Println(userId)
	rawString := fmt.Sprintf("SELECT * FROM todolists WHERE user_id = %v", userId)
	resultSql, errorQuery := t.db.Query(rawString)

	if errorQuery != nil {
		fmt.Println("Error !!", errorQuery)
		c.JSON(500, errorQuery)
	}
	defer resultSql.Close()

	var todos []Todolist
	for resultSql.Next() {
		var todo Todolist
		err := resultSql.Scan(&todo.Id, &todo.Status, &todo.Description, &todo.Title, &todo.UserId)

		if err != nil {
			panic(err)
		}
		todos = append(todos, todo)
	}
	c.JSON(200, todos)
}
