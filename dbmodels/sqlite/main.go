package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate to schema
	db.AutoMigrate(&Product{})

	// Create
	product001 := Product{Code: "PDOO1", Price: 100}
	db.Create(&product001)
	time.Sleep(time.Second)

	// Read
	var result Product
	db.First(&result, product001.ID) // find product with integer primary key
	fmt.Println(result)
	fmt.Printf("Read first ID ID: %d, Code: %s, Price: %d\n", result.ID, result.Code, result.Price)
	db.First(&result, "code = ?", product001.Code) // find product with code 001
	fmt.Printf("Read specified Code ID: %d, Code: %s, Price: %d\n", result.ID, result.Code, result.Price)

	// Update - update a single field, e.g., product's price to 200
	db.Model(&result).Update("price", 300)
	// Read

	db.First(&result, "code = ?", product001.Code) // find product with code PD001
	fmt.Printf("Read on single update ID: %d, Code: %s, Price: %d\n", result.ID, result.Code, result.Price)

	// Update

	//	Update multiple fields using struct
	db.Model(&result).Updates(Product{Price: 300, Code: "F42"}) // non-zero fields
	// Read
	fmt.Println("Result Code: " + result.Code)

	db.First(&result, "code = ?", result.Code) // find product with code F42
	fmt.Printf("Read on update ID: %d, Code: %s, Price: %d\n", result.ID, result.Code, result.Price)

	//	Update multiple fields using map
	db.Model(&result).Updates(map[string]interface{}{"price": 100, "code": "P001"})
	// Read
	db.First(&result, "code = ?", result.Code) // find product with code 001
	fmt.Printf("Read on update ID: %d, Code: %s, Price: %d\n", result.ID, result.Code, result.Price)

	// Delete - delete product
	db.Delete(&result, result.ID)                                                         // delete product with integer primary key
	db.Unscoped().Where("deleted_at IS NOT NULL").First(&result, "code = ?", result.Code) // find product with code 001
	fmt.Printf("Read on delete ID: %d, Code: %s, Price: %d, Deleted At: %v\n",
		result.ID, result.Code, result.Price, result.DeletedAt)

	db.Unscoped().Delete(&result) // find product with code 001
	fmt.Printf("Deleted product with ID: %d, Code: %s, Price: %d, Deleted At: %v\n",
		result.ID, result.Code, result.Price, result.DeletedAt)

	var results []Product
	db.Find(&results)
	fmt.Println(results)

	var deletedProducts []Product
	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&deletedProducts)
	fmt.Println(deletedProducts)

	for _, product := range deletedProducts {
		db.Unscoped().Model(&product).UpdateColumn("deleted_at", nil)
	}

}
