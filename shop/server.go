package main

import (
	"fmt"
	"os"

	"go-getting-started/shop/pkgs/apis"
	"go-getting-started/shop/pkgs/database/gorm"
	"go-getting-started/shop/pkgs/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, _ := gorm.InitializeDB()
	db.AutoMigrate(&models.Product{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.CartItem{})
	app := fiber.New()
	port := ":3000"
	args := os.Args[1:]
	fmt.Println(args)
	if len(args) > 0 {
		port = fmt.Sprintf(":%s", args[0])
	}

	v1 := app.Group("/v1")
	apis.RegisterApiRoutes(v1, db)

	// v2 := app.Group("/v2")
	// apis.RegisterApiRoutes(v2, db)

	err := app.Listen(port)
	if err != nil {
		panic(err)
	}
}
