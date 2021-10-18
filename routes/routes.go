package routes

import (
	"cash-app/constant"
	"cash-app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	//--------------------User Controllers--------------------//
	e.POST("/users/signup", controllers.UserSignUpController) // user SignUp
	e.POST("/users/signin", controllers.UserSignInController) // user SignIn

	//--------------------Authorization--------------------//
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	//--------------------Controllers User--------------------//
	eJwt.GET("/users", controllers.GetAllUsersController)            // Get all users data
	eJwt.GET("/users/:user_id", controllers.GetUserByIdController)   // Get user data
	eJwt.PUT("/users/:user_id", controllers.UpdateUserController)    // Update user data
	eJwt.DELETE("/users/:user_id", controllers.DeleteUserController) // Delete user data

}
