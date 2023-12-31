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
	r.Run()
}
