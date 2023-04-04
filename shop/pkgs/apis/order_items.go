package apis

import (
	"go-getting-started/shop/pkgs/models"

	"github.com/gofiber/fiber/v2"
)

// ------------------------OrderItem Handlers------------------------------//
type OrderItemRequestBody struct {
	CartNumber string
	Quantity   int
}

func (h Handler) GetOrderItems(c *fiber.Ctx) error {
	var orderItems []models.OrderItem

	err := h.DB.Find(&orderItems).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "All OrderItems",
		"data":    orderItems,
	})
}
