package main

import (
	_ "database/sql"
	"fmt"
	"time"

	//_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres" // to customize drivers

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=localhost user=test_user password=test_pass dbname=test_db port=5432 sslmode=disable Timezone=Asia/Shanghai"
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
	//product001 := &Product{Code: "PDOO1", Price: 100}
	db.Create(&Product{Code: "PDOO1", Price: 100})
	time.Sleep(time.Second)

	// Read
	product001 := &Product{}
	db.First(&product001, 0) // find product with integer primary key
	fmt.Printf("Read first ID ID: %d, Code: %s, Price: %d", product001.ID, product001.Code, product001.Price)
	db.First(&product001, "code = ?", "PD001") // find product with code 001
	fmt.Printf("Read specified Code ID: %d, Code: %s, Price: %d", product001.ID, product001.Code, product001.Price)

	// Update - update a single field, e.g., product's price to 200
	db.Model(&product001).Update("Price", 200)
	// Read
	db.First(&product001, "code = ?", "PD001") // find product with code 001
	fmt.Printf("Read on single update ID: %d, Code: %s, Price: %d", product001.ID, product001.Code, product001.Price)

	// Update

	//	Update multiple fields using struct
	db.Model(&product001).Updates(Product{Price: 300, Code: "F42"}) // non-zero fields
	// Read
	db.First(&product001, "code = ?", "PD001") // find product with code 001
	fmt.Printf("Read on update ID: %d, Code: %s, Price: %d", product001.ID, product001.Code, product001.Price)

	//	Update multiple fields using map
	db.Model(&product001).Updates(map[string]interface{}{"Price": 100, "Code": "P001"})
	// Read
	db.First(&product001, "code = ?", "PD001") // find product with code 001
	fmt.Printf("Read on update ID: %d, Code: %s, Price: %d", product001.ID, product001.Code, product001.Price)

	// Delete - delete product
	db.Delete(&product001, 1)                  // delete product with integer primary key
	db.First(&product001, "code = ?", "PD001") // find product with code 001
	fmt.Printf("Read on delete ID: %d, Code: %s, Price: %d", product001.ID, product001.Code, product001.Price)
}
