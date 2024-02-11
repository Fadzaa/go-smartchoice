package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"smartchoice/domain/candidate_pair"
	"smartchoice/domain/candidate_profile"
	"smartchoice/domain/debate"
	"smartchoice/domain/party"
)

func ConnectToDatabase() *gorm.DB {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&candidate_pair.CandidatePair{},
		&candidate_profile.CandidateProfile{},
		&party.Party{},
		&candidate_profile.Achievement{},
		&debate.Debate{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected")

	return db
}
