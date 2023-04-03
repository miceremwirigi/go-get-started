package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"go-getting-started/shop/pkgs/apis"
	"go-getting-started/shop/pkgs/models"
	"go-getting-started/shop/pkgs/database/gorm"
)

func main() {
	db, _ := gorm.InitializeDB()
	db.AutoMigrate(&models.Product{})
	app := fiber.New()
	port := ":3000"
	args := os.Args[1:]
	fmt.Println(args)
	if len(args) > 0 {
		port = fmt.Sprintf(":%s",args[0])
	}

	h := apis.Handler{DB: db}
	app.Get("/products", h.GetProducts)
	app.Post("/products", h.AddProduct)
	app.Put("/products/:id", h.UpdateProduct)
	app.Get("/products/:id", h.GetProduct)
	app.Delete("/products/:id", h.DeleteProduct)

	err := app.Listen(port)
	if err != nil {
		panic(err)
	}
}
