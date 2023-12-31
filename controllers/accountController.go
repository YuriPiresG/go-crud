package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuripireg/go-crud/initializers"
	"github.com/yuripireg/go-crud/models"
)

func AccountCreate(c *gin.Context) {

	var body struct {
		Body  string
		AccountNumber string
		Balance float64
	}

	c.Bind(&body)

	account := models.Account{AccountNumber: "1", Balance: 1000}

	result := initializers.DB.Create(&account)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"account": account,
	})
}
