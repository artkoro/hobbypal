package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() *gorm.DB {
	connectionString := "postgres://artem:zen@localhost:5432/hobbypal"
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		panic(err)
	}
	return db
}
