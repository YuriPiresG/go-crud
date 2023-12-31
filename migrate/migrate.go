package main

import (
	"github.com/yuripireg/go-crud/initializers"
	"github.com/yuripireg/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Account{})
	initializers.DB.AutoMigrate(&models.Payments{})
}
