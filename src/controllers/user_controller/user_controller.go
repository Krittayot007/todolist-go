package user_controller

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUser(*gin.Context)
	CreateUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userController struct {
	db *sql.DB
}

type User struct {
	Id 	 		int 	`json:"id"`
	Name 		string 	`json:"name"`
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}

func NewUserController(db *sql.DB) UserController { // connect to database
	return userController{db: db}
}

func (u userController) GetAllUser(res *gin.Context) {
	result, err := u.db.Query("SELECT id,name,email FROM users")
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	defer result.Close()
	
	var users []User
	for result.Next() {
		var user User 
		err := result.Scan(&user.Id, &user.Name, &user.Email) // ส่ง address ไปเช็คว่ามีหรือไม่

		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	res.JSON(200, users)
}

func (u userController) CreateUser(c *gin.Context) {
	var dataUserCreate User
	errorRequest := c.BindJSON(&dataUserCreate)
	fmt.Printf("data from req.body %+v\n", errorRequest)

	if errorRequest != nil {
		fmt.Printf("Error %+v", errorRequest)
	}

	create, _ := u.db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	createResult, createError := create.Exec(dataUserCreate.Name, dataUserCreate.Email, dataUserCreate.Password)
	fmt.Println("result of creation", createResult)

	if createError != nil {
		fmt.Println(createError)
		c.JSON(500, createError)
	}
	defer create.Close()

	c.JSON(201, dataUserCreate)
}

func (u userController) UpdateUser(c *gin.Context) {
	params := c.Params
	id := params[0].Value

	dataUpdate := make(map[string]string)
	errorBind := c.BindJSON(&dataUpdate)
	fmt.Println(errorBind)

	key := make([]string, 0, 2)
	value := make([]interface{}, 0, 3)

	for k,v := range dataUpdate {
		key = append(key, k)
		value = append(value, v)
	}
	value = append(value, id)

	keyString := strings.Join(key, " = ?, ") // จะคั่นให้ระหว่่าง struct
	rawString := fmt.Sprintf("UPDATE users SET %s = ? WHERE id = ? ", keyString)

	prepareUpdate, _ := u.db.Prepare(rawString)
	resultUpdate, updateError := prepareUpdate.Exec(value...)
	fmt.Println(resultUpdate, updateError)
	defer prepareUpdate.Close()

	if updateError != nil {
		c.JSON(500, updateError)
	}
	
	c.JSON(200, resultUpdate)
}

func (u userController) DeleteUser(d *gin.Context) {
	params := d.Params
	id := params[0].Value

	prepare, _ := u.db.Prepare("DELETE FROM users where id = ?")
	_, deleteError := prepare.Exec(id)
	if deleteError != nil {
		fmt.Println(deleteError.Error())
		d.JSON(500, deleteError)
	}

	d.JSON(200, "delete" + id + "success")
} 