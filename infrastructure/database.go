package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"smartchoice/domain/party"
)

func ConnectToDatabase() *gorm.DB {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&party.Party{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")

	return db
}
