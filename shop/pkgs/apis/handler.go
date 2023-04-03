package apis

import (
	"fmt"
	"go-getting-started/shop/pkgs/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

type ProductRequestBody struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Price string `json:"price"`
	HirePurchase bool `json:"hire_purchase"`
}

func (h Handler) GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	err := h.DB.Find(&products).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}

func (h Handler) GetProduct(c *fiber.Ctx) error {
	var product models.Product

	id := c.Params("id")
	err := h.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&product)
}

func (h Handler) AddProduct(c *fiber.Ctx) error {
	var product models.Product
	var body ProductRequestBody
	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	product.Name = body.Name
	product.Code = body.Code
	product.Price = body.Price
	product.HirePurchase = body.HirePurchase

	err = h.DB.Create(&product).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(&product)
}

func (h Handler) UpdateProduct( c *fiber.Ctx) error {
	var product models.Product
	var body ProductRequestBody
	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id := c.Params("id")
	err = h.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	product.Name = body.Name
	product.Code = body.Code
	product.Price = body.Price
	product.HirePurchase = body.HirePurchase
	
	err = h.DB.Save(&product).Error
	if err != nil {
		return fiber.NewError(fiber.StatusExpectationFailed, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&product)
}

func (h Handler) DeleteProduct( c *fiber.Ctx) error {
	var product models.Product

	id := c.Params("id")
	err := h.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	err = h.DB.Delete(&product).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotModified, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Product of id %s deleted", id),
	})
}

func RegisterApiRoutes(url fiber.Router, db *gorm.DB) {
	h := Handler{DB: db}
	productsRoutes := url.Group("/products")
	
	productsRoutes.Get("/", h.GetProducts)
}
