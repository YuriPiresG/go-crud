package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Balance       float64
	Payments	  []Payments
}
