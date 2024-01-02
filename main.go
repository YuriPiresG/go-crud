package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuripireg/go-crud/controllers"
	"github.com/yuripireg/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/account", controllers.AccountCreate)
	r.GET("/account", controllers.AccountsGet)
	r.GET("/account/:id", controllers.AccountGet)
	r.PATCH("/account/:id", controllers.AccountUpdate)
	r.DELETE("/account/:id", controllers.AccountDelete)
	r.Run()
}
