package model

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB = GoConnect()

func GoConnect() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	mySQLHost := os.Getenv("DB_HOST")
	mySQLUser := os.Getenv("DB_USER")
	mySQLPass := os.Getenv("DB_PASS")
	mySQLDBName := os.Getenv("DB_NAME")
	mySQLDBPort := os.Getenv("DB_PORT")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mySQLUser,
		mySQLPass,
		mySQLHost,
		mySQLDBPort,
		mySQLDBName,
	)
	fmt.Println("BLOG dns:", dns)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to Gorm database" + err.Error())
	}
	//migration of table structure
	return db
}

type Date struct {
	Created time.Time `gorm:"autoCreateTime"`
	Updated time.Time `gorm:"autoUpdateTime"`
}
