package db

import (
	"fmt"
	"log"

	"github.com/naphat/gobapi/cmd/api/internal/core/domain"
	"github.com/naphat/gobapi/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect เชื่อมต่อฐานข้อมูล PostgreSQL
func Connect(config *config.Config) *gorm.DB {

	// สร้าง connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBSSL,
	)

	// เชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อฐานข้อมูลได้: %v", err)
	}

	// สร้างตารางในฐานข้อมูลจากโมเดล
	migrateDatabase(db)

	log.Println("✅ เชื่อมต่อฐานข้อมูล PostgreSQL สำเร็จ")

	return db
}

// migrateDatabase สร้างตารางในฐานข้อมูลจากโมเดล
func migrateDatabase(db *gorm.DB) {
	// Auto migrate สร้างตารางในฐานข้อมูลจากโมเดล
	err := db.AutoMigrate(
		&domain.Role{},
		&domain.Permission{},
		&domain.User{},
		&domain.Category{},
		&domain.Product{},
		&domain.ProductImage{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Transaction{},
	)

	if err != nil {
		log.Fatalf("ไม่สามารถสร้างตารางในฐานข้อมูลได้: %v", err)
	}

	log.Println("✅ สร้างตารางในฐานข้อมูลสำเร็จ")
}