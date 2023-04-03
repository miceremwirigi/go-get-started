package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=localhost user=test_user password=test_pass dbname=test_db port=5432 sslmode=disable Timezone=Africa/Nairobi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	/*
		// We are using pgx as postgres’s database/sql driver,
		// it enables prepared statement cache by default
		// to reduce the need to compile repeated executable SQL statements, to disable it:

		// -----------------Optional-----------------//
		// --------Disable prepared statement cache---------//
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})

		// -----------------Optional-----------------//
		// -------------Customize Driver------------ //

		db, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: "cloudsqlpostgress",
			DSN:        "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
		}), &gorm.Config{})


		// -----------------Optional-----------------//

		// -------------Existing Database Connection------------ //

		sqlDB, err := sql.Open("pgx", "mydb_dsn")
		if err != nil {
			panic("failed to find existing database")
		}
		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db = gormDB
	*/

	// *********Working with the Database Connection******//
	// Migrate to schema
	db.AutoMigrate(&Product{})

	// Create
	var product Product = Product{Code: "P001", Price: 150}
	db.Create(&product)

	// Read
	result := Product{}
	var allProducts []Product
	db.Find(&allProducts)
	fmt.Println("All products: \n", allProducts)
	db.First(&result, "Code = ?", product.Code) // Find first product whose code is P001

	// Update - update a single field, e.g., product's price to 200
	db.Model(&result).Update("Price", 200)
	// Read
	db.First(&result, "code = ?", product.Code) // find product with code 001
	fmt.Printf("Read on single update ID: %d, Code: %s, Price: %d\n", result.ID, result.Code, result.Price)

	//	Update multiple fields using struct
	db.Model(&result).Updates(Product{Price: 300, Code: "F42"}) // non-zero fields
	// Read
	db.First(&result, "code = ?", "F42") // find product with code 001
	fmt.Printf("Read on update ID: %d, Code: %s, Price: %d\n", result.ID, result.Code, result.Price)

	//	Update multiple fields using map
	db.Model(&result).Updates(map[string]interface{}{"Price": 100, "Code": "P001"})
	// Read
	db.First(&result, "code = ?", product.Code) // find product with code 001
	fmt.Printf("Read on update ID: %d, Code: %s, Price: %d\n", result.ID, result.Code, result.Price)

	// Delete - delete product
	db.Delete(&result, 1)                       // delete product with integer primary key
	db.First(&result, "code = ?", product.Code) // find product with code 001
	fmt.Printf("Read on delete ID: %d, Code: %s, Price: %d\n", result.ID, result.Code, result.Price)

}
