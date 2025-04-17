package main

import (
	"fmt"
	"log"
	"os"

	"github.com/neboman11/music-wishlist-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func open_database() {
	backend := os.Getenv("DATABASE_BACKEND")
	if backend == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_DB"))
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %s", err)
		}
	} else if backend == "mariadb" || backend == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_DB"))
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %s", err)
		}
	} else {
		log.Fatal("Unsupported database backend")
	}

	db.AutoMigrate(&models.Want{})
}
