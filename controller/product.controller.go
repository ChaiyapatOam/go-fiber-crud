package controller

import (
	"fmt"

	"github.com/chaiyapatoam/go-fiber-crud/database"
	"github.com/chaiyapatoam/go-fiber-crud/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	// define product lsit
	var products []model.Product
	result := database.DB.Find(&products)
	fmt.Println(result, products)
	c.JSON(products)
	return nil
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	p := new(model.Product)
	database.DB.Where("id= ?", id).First(&p)
	c.JSON(p)
	return nil
}

func AddProduct(c *fiber.Ctx) error {
	p := new(model.Product)
	if err := c.BodyParser(p); err != nil {
		fmt.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	res := database.DB.Create(&p)
	if res.Error != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": res.Error,
		})
	}
	type Response struct {
		Success bool
		Message string
	}
	data := Response{
		Success: true,
		Message: "Created!",
	}
	return c.Status(201).JSON(data)
}

func EditProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	p := new(model.Product)
	c.BodyParser(p)
	database.DB.Where("id= ?", id).Update(p.Name, p.Price)
	return nil
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	p := new(model.Product)
	database.DB.Where("id= ?", id).Delete(&p)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Deleted!",
	})
	return nil
}
