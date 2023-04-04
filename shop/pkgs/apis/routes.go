package apis

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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
	cartItemsRoutes := url.Group("/cart_items")
	cartItemsRoutes.Get("/", h.GetCartItems)
	cartItemsRoutes.Post("/", h.AddCartItem)
	cartItemsRoutes.Put("/:id", h.UpdateCartItem)
	cartItemsRoutes.Get("/:id", h.GetCartItem)
	cartItemsRoutes.Delete("/:id", h.DeleteCartItem)

	//------------OrderRoutes--------------------//
	ordersRoutes := url.Group("/orders")
	ordersRoutes.Get("/", h.GetOrders)
	// ordersRoutes.Post("/", h.AddOrder)
	// ordersRoutes.Put("/:id", h.UpdateOrder)
	// ordersRoutes.Get("/:id", h.GetOrder)
	// ordersRoutes.Delete("/:id", h.DeleteOrder)

	//------------OrderItemsRoutes--------------------//
	orderItemsRoutes := url.Group("/order_items")
	orderItemsRoutes.Get("/", h.GetOrderItems)
	// orderItemsRoutes.Post("/", h.AddOrderItem)
	// orderItemsRoutes.Put("/:id", h.UpdateOrderItem)
	// orderItemsRoutes.Get("/:id", h.GetOrderItem)
	// orderItemsRoutes.Delete("/:id", h.DeleteOrderItem)

}
