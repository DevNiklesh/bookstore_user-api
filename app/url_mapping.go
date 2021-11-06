package app

import (
	pingControllers "github.com/DevNiklesh/bookstore_user-api/controllers/ping_controllers"
	usersControllers "github.com/DevNiklesh/bookstore_user-api/controllers/users_controllers"
)

func mapUrls() {
	router.GET("/ping", pingControllers.Ping)

	router.GET("/user/:id", usersControllers.GetUser)
	router.POST("/users", usersControllers.CreateUser)
}
