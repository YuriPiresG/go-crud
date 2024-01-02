package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuripireg/go-crud/services"
)

func AccountCreate(c *gin.Context) {
	var body struct {
		Balance float64
	}

	c.Bind(&body)

	account, err := services.CreateAccount(body.Balance)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"account": account,
	})
}

func AccountGet(c *gin.Context) {
	id := c.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	account, err := services.GetAccount(uint(idUint))
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"account": account,
	})
}

func AccountsGet(c *gin.Context) {
	accounts, err := services.GetAccounts()
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"accounts": accounts,
	})
}

func AccountUpdate(c *gin.Context) {
	var body struct {
		Balance float64
	}

	c.Bind(&body)

	id := c.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}

	account, err := services.UpdateAccount(uint(idUint), body.Balance)
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"account": account,
	})
}

func AccountDelete(c *gin.Context) {
	id := c.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.Status(400)
		return
	}

	err = services.DeleteAccount(uint(idUint))
	if err != nil {
		c.Status(400)
		return
	}
	c.Status(200)

}
