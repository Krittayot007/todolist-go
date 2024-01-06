package todo_controller

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	GetAllTodoWithUserId(*gin.Context)
	CreateTodo(*gin.Context)
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
		fmt.Println("Error !!", errorQuery.Error())
		c.JSON(500, errorQuery.Error())
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

func (t todoController) CreateTodo(c *gin.Context) {
	params := c.Params
	userId := params[0].Value

	var dataCreate Todolist
	errorRequest := c.BindJSON(&dataCreate)
	fmt.Printf("data from req.body %+v", errorRequest)

	if errorRequest != nil {
		fmt.Printf("Error !! %+v", errorRequest.Error())
	}

	rawString := fmt.Sprintf("INSERT INTO todolists (status, description, title, user_id) VALUES (?, ?, ?, ?)")
	create, _ := t.db.Prepare(rawString)
	createResult, createError := create.Exec(&dataCreate.Status, &dataCreate.Description, &dataCreate.Title, &userId)

	if createError != nil {
		fmt.Println("Error!!", createError.Error())
		c.JSON(500, createError.Error())
	}
	defer create.Close()

	c.JSON(201, createResult)
	// token := c.GetHeader("Authorization") get header for token

}
