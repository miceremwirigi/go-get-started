package apis

import (
	"go-getting-started/shop/pkgs/models"

	"github.com/gofiber/fiber/v2"
)

// ------------------------Order Handlers------------------------------//
type OrderRequestBody struct {
	CartNumber string
	Quantity   int
}

func (h Handler) GetOrders(c *fiber.Ctx) error {
	var orders []models.CartItem

	err := h.DB.Find(&orders).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "All Orders",
		"data":    orders,
	})
}
