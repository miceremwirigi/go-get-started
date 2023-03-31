package main

import (
	_ "gorm.io/driver/sqlserver"
	_ "gorm.io/gorm"
)

// func main() {
// 	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
// 	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
// }
