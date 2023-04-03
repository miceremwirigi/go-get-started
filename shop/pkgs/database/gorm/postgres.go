package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (db *gorm.DB, err error) {
	dsn := "host=localhost user=test_user password=test_pass dbname=test_db port=5432 sslmode=disable Timezone=Africa/Nairobi"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
