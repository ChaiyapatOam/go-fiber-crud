package controller

import (
	"fmt"

	"github.com/chaiyapatoam/go-fiber-crud/database"
	"github.com/chaiyapatoam/go-fiber-crud/model"
	"github.com/gofiber/fiber/v2"
)
func GetAllProducts(c *fiber.Ctx) error {
	products := new(model.Products)
	result := database.DB.Find(&products)
	fmt.Println(result,products)
	c.JSON(products)
	 return nil
}

func GetProductById(c *fiber.Ctx) error {
	
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
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": res.Error,
		  })
		return nil
	}
	 return nil
}

func EditProduct(c *fiber.Ctx) error {
	
	 return nil
}

func DeleteProduct(c *fiber.Ctx) error {
	
	 return nil
}