package routers

import (
	"todolist-go/src/controllers/todo_controller"
	user_controller "todolist-go/src/controllers/user_controller"
	"todolist-go/src/models"
	"todolist-go/src/repository"
)

func UserRouter(server models.Server, PATH string) {
	userRouter := server.Engine.Group(PATH)
	db := repository.Database()

	userController := user_controller.NewUserController(db)

	userRouter.GET("/getAllUser", userController.GetAllUser)
	userRouter.POST("/createUser", userController.CreateUser)
	userRouter.PATCH("/updateUser/:id", userController.UpdateUser)
	userRouter.DELETE("/deleteUser/:id", userController.DeleteUser)
}

func TodoRouter(server models.Server, PATH string) {
	todoRouter := server.Engine.Group(PATH)
	db := repository.Database()

	todoController := todo_controller.NewTodoController(db)

	todoRouter.GET("/getTodoWithUserId", todoController.GetAllTodoWithUserId)
	todoRouter.POST("/createTodo/:userId", todoController.CreateTodo)
}
