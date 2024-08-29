package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/phuslu/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Err(errors.New("FAILED TO INITIALIZE DATABASE : " + err.Error())).Msg("")
		os.Exit(1)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	log.Info().Msg("=== DATABASE CONNECTION SUCCESSFULLY ===")
	return db
}
