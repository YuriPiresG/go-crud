package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	AccountNumber string
	Balance       float64
	Payments	  []Payments
}
