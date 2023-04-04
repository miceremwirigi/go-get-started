package apis

import (
	"fmt"
	"go-getting-started/shop/pkgs/models"

	"github.com/gofiber/fiber/v2"
)

type CartRequestBody struct {
	CartNumber string `json:"cart_number"`
	Customer   string `json:"customer"`
}

func (h Handler) GetCarts(c *fiber.Ctx) error {
	var carts []models.Cart

	err := h.DB.Find(&carts).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "All carts",
		"data":    carts,
	})
}

func (h Handler) GetCart(c *fiber.Ctx) error {
	var cart models.Cart

	id := c.Params("id")
	err := h.DB.First(&cart, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Showing Cart ID, %s", id),
		"data":    cart,
	})
}

func (h Handler) AddCart(c *fiber.Ctx) error {
	var cart models.Cart
	var body CartRequestBody

	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cart.CartNumber = body.CartNumber
	cart.Customer = body.Customer

	err = h.DB.Create(&cart).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotImplemented, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Added Cart of Cart number, %s to database", body.CartNumber),
		"data":    cart,
	})
}

func (h Handler) UpdateCart(c *fiber.Ctx) error {
	var cart models.Cart
	var body CartRequestBody

	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id := c.Params("id")
	err = h.DB.First(&cart, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	cart.CartNumber = body.CartNumber
	cart.Customer = body.Customer

	err = h.DB.Save(&cart).Error
	if err != nil {
		return fiber.NewError(fiber.StatusExpectationFailed, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Updated Cart of ID, %s to database", id),
		"data":    cart,
	})
}

func (h Handler) DeleteCart(c *fiber.Ctx) error {
	var cart models.Cart

	id := c.Params("id")

	err := h.DB.First(&cart, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	err = h.DB.Delete(&cart, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotModified, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Cart of id %s deleted", id),
	})
}
