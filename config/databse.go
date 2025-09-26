package config

import (
	"fmt"
	"log"

	"github.com/naol86/go/authGo/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConfig(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBUser,
		env.DBPassword,
		env.DBHost,
		env.DBPort,
		env.DBName,
	)
	fmt.Println("Connecting to database with DSN:", dsn) // Debugging line

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
		panic(err)
	}
	// Auto-migrate the User model
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}
	return db
}

func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
		panic(err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("failed to close database connection: %v", err)
		panic(err)
	}
	log.Println("Database connection closed")
}
