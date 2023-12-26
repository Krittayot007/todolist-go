package routers

import (
	user_controller "todolist-go/src/controllers/user_controller"
	"todolist-go/src/models"
	"todolist-go/src/repository"
)

func UserRouter(server models.Server, PATH string) {
	userRouter := server.Engine.Group(PATH)
	db := repository.Database()

	usercontroller := user_controller.NewUserController(db) 

	userRouter.GET("/getAllUser", usercontroller.GetAllUser)
	userRouter.POST("/createUser", usercontroller.CreateUser)
	userRouter.PATCH("/updateUser/:id", usercontroller.UpdateUser)
	userRouter.DELETE("/deleteUser/:id", usercontroller.DeleteUser)
}