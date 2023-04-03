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

// ------------------------Product Handlers------------------------------//

type ProductRequestBody struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	Price        string `json:"price"`
	HirePurchase bool   `json:"hire_purchase"`
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

func (h Handler) UpdateProduct(c *fiber.Ctx) error {
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

func (h Handler) DeleteProduct(c *fiber.Ctx) error {
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

// ------------------------Cart Handlers------------------------------//
type CartRequestBody struct {
	CartNumber string
	Quantity   int
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
	cart.Quantity = body.Quantity

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
	cart.Quantity = body.Quantity

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

// ------------------------CartItem Handlers------------------------------//
type CartItemRequestBody struct {
	CartNumber string
	Quantity   int
}

func (h Handler) GetCartItems(c *fiber.Ctx) error {
	var cartItems []models.CartItem

	err := h.DB.Find(&cartItems).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "All CartItems",
		"data":    cartItems,
	})
}

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

//----------------Register Routes-------------------//

func RegisterApiRoutes(url fiber.Router, db *gorm.DB) {
	h := Handler{DB: db}

	//------------ProductsRoutes--------------------//
	productsRoutes := url.Group("/products")
	productsRoutes.Get("/", h.GetProducts)
	productsRoutes.Post("/", h.AddProduct)
	productsRoutes.Put("/:id", h.UpdateProduct)
	productsRoutes.Get("/:id", h.GetProduct)
	productsRoutes.Delete("/:id", h.DeleteProduct)

	//-----------------CartRoutes--------------------//
	cartsRoutes := url.Group("/carts")
	cartsRoutes.Get("/", h.GetCarts)
	cartsRoutes.Post("/", h.AddCart)
	cartsRoutes.Put("/:id", h.UpdateCart)
	cartsRoutes.Get("/:id", h.GetCart)
	cartsRoutes.Delete("/:id", h.DeleteCart)

	//-----------------CartItemsRoutes--------------------//
	cartItemsRoutes := url.Group("/cartitems")
	cartItemsRoutes.Get("/", h.GetCartItems)
	//  cartItemsRoutes.Post("/", h.AddCartItem)
	//  cartItemsRoutes.Put("/:id", h.UpdateCartItem)
	//  cartItemsRoutes.Get("/:id", h.GetCartItem)
	// cartItemsRoutes.Delete("/:id", h.DeleteCartItem)

	//------------OrderRoutes--------------------//
	ordersRoutes := url.Group("/orders")
	ordersRoutes.Get("/", h.GetOrders)
	// ordersRoutes.Post("/", h.AddOrder)
	// ordersRoutes.Put("/:id", h.UpdateOrder)
	// ordersRoutes.Get("/:id", h.GetOrder)
	// ordersRoutes.Delete("/:id", h.DeleteOrder)

	//------------OrderItemsRoutes--------------------//
	orderItemsRoutes := url.Group("/orderItems")
	orderItemsRoutes.Get("/", h.GetOrderItems)
	// orderItemsRoutes.Post("/", h.AddOrderItem)
	// orderItemsRoutes.Put("/:id", h.UpdateOrderItem)
	// orderItemsRoutes.Get("/:id", h.GetOrderItem)
	// orderItemsRoutes.Delete("/:id", h.DeleteOrderItem)

}
