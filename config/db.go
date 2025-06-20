package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {

	dsn := os.Getenv("DB_CONNECT_STR")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Kết nối DB thất bại:", err)
	}
	log.Println("✅ Kết nối DB thành công")
	DB = db
}
