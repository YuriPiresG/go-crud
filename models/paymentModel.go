package models

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	AccountID uint
	Value           float64
	PaymentDate     string
}
