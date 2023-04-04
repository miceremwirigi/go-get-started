package apis

import (
	"fmt"
	"go-getting-started/shop/pkgs/models"

	"github.com/gofiber/fiber/v2"
)

type CartItemRequestBody struct {
	ItemName  string `json:"item_name"`
	CartID    int    `json:"cart_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (h Handler) GetCartItems(c *fiber.Ctx) error {
	var cartItems []models.CartItem

	err := h.DB.Preload("Cart").Preload("Product").Find(&cartItems).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "All CartItems",
		"data":    cartItems,
	})
}

func (h Handler) GetCartItem(c *fiber.Ctx) error {
	var cartItem models.Cart

	id := c.Params("id")
	err := h.DB.First(&cartItem, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Showing CartItem ID, %s", id),
		"data":    cartItem,
	})
}

func (h Handler) AddCartItem(c *fiber.Ctx) error {
	var cartItem models.CartItem
	var body CartItemRequestBody

	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cartItem.ItemName = body.ItemName
	cartItem.CartID = body.CartID
	cartItem.ProductID = body.ProductID
	cartItem.Quantity = body.Quantity

	err = h.DB.Create(&cartItem).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotImplemented, err.Error())
	}

	var result models.CartItem
	err = h.DB.Preload("Cart").Preload("Product").First(&result, cartItem.ID).Error
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Added %d CartItems of Product %s to database", body.Quantity, cartItem.ItemName),
		"data":    cartItem,
	})
}

func (h Handler) UpdateCartItem(c *fiber.Ctx) error {
	var cartItem models.CartItem
	var body CartItemRequestBody

	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id := c.Params("id")
	err = h.DB.First(&cartItem, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	cartItem.ItemName = body.ItemName
	cartItem.ProductID = body.ProductID
	cartItem.CartID = body.CartID
	cartItem.Quantity = body.Quantity

	err = h.DB.Save(&cartItem).Error
	if err != nil {
		return fiber.NewError(fiber.StatusExpectationFailed, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Updated %s CartItem of CartID %d to database", cartItem.Product.Name, cartItem.Cart.ID),
		"data":    cartItem,
	})
}

func (h Handler) DeleteCartItem(c *fiber.Ctx) error {
	var cartItem models.CartItem

	id := c.Params("id")

	err := h.DB.First(&cartItem, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	err = h.DB.Delete(&cartItem, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotModified, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Deleted %s CartItem of CartNumber %s to database", cartItem.Product.Name, cartItem.Cart.CartNumber),
	})
}
