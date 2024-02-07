package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectToDatabase() *gorm.DB {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//err = db.AutoMigrate(&auth.User{}, &job.Job{}, &party.Course{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")

	return db
}
