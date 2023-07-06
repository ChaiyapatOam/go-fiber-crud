package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type Products struct {
	Products []Product `json:"products"`
}