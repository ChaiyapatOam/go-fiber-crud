package controller
import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetCart(c *fiber.Ctx) error{
	fmt.Println("getCart")
	return nil
}