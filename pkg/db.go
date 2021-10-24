package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	DB *gorm.DB
}

type DatabaseOption struct {
	Host string
	Port int
	User string
	Pass string
	Db   string
}

func NewMysql(option DatabaseOption) *Database {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", option.User, option.Pass, option.Host, option.Port, option.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	return &Database{
		DB: db,
	}

}
