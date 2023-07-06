package database

import (
	"fmt"

	"github.com/chaiyapatoam/go-fiber-crud/config"
	"github.com/chaiyapatoam/go-fiber-crud/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connect() error {
	var err error
	dsn := config.Config("DSN")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	err = DB.AutoMigrate(&model.Product{})
	if err != nil {
		panic("migration fail")
	}
	
	insertProduct := &model.Product{Name: "Second", Price: 120}

	DB.Create(insertProduct)
	fmt.Println("Connected DB")
	return nil
}