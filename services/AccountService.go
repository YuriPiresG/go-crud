package services

import (
	"github.com/yuripireg/go-crud/initializers"
	"github.com/yuripireg/go-crud/models"
)

func CreateAccount(balance float64) (models.Account, error) {
	account := models.Account{Balance: balance}

	result := initializers.DB.Create(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

func GetAccount(id uint) (models.Account, error) {
	var account models.Account

	result := initializers.DB.First(&account, id)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

func GetAccounts() ([]models.Account, error) {
	var accounts []models.Account

	result := initializers.DB.Find(&accounts)

	if result.Error != nil {
		return []models.Account{}, result.Error
	}

	return accounts, nil
}

func UpdateAccount(id uint, balance float64) (models.Account, error) {
	var account models.Account

	result := initializers.DB.First(&account, id)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	account.Balance = balance

	result = initializers.DB.Save(&account)

	if result.Error != nil {
		return models.Account{}, result.Error
	}

	return account, nil
}

func DeleteAccount(id uint) error {
	var account models.Account

	result := initializers.DB.First(&account, id)

	if result.Error != nil {
		return result.Error
	}

	result = initializers.DB.Delete(&account)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
