package main

import (
	_ "database/sql"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

// func main() {
// 	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	// ------------------Optional-------------------//
// 	// -----------Advanced configations------------ //

// 	db, err = gorm.Open(mysql.New(mysql.Config{
// 		DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?chrset=utf8mb4&parseTime=True&loc=Local", // data source name
// 		DefaultStringSize:         256,                                                                          // default size for string fields
// 		DisableDatetimePrecision:  true,                                                                         // disable datetimeprecision, which not supported before MySQL 5.6
// 		DontSupportRenameIndex:    true,                                                                         // drop & create when rename index, rename index not supported before MySQL 5.7
// 		DontSupportRenameColumn:   true,                                                                         // `change` when rename column, rename column could not supported before MySQL 8
// 		SkipInitializeWithVersion: false,                                                                        // auto configure based on the current MySQL version
// 	}), &gorm.Config{})

// 	// -----------------Optional-----------------//
// 	// -------------Customize Driver------------ //
// 	db, err = gorm.Open(mysql.New(mysql.Config{
// 		DriverName: "my_mysql_driver",
// 		DSN:        "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local", // data ource name
// 	}), &gorm.Config{})

// 	// --------Existing Database Connection --------//
// 	sqlDB, err := sql.Open("mysql", "mydb_dsn")
// 	gormDb, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn: sqlDB,
// 	}), &gorm.Config{})

// }
