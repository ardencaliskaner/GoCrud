package db

import (
	"fmt"
	"log"
	"os"

	"GoCrud/pkg/db/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func setupDB(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}

func ConnectDB() *gorm.DB {

	dbHost := os.Getenv("POSTGRES_HOST")

	fmt.Println(dbHost)

	if dbHost == "" {
		err := godotenv.Load("local.env")
		dbHost = os.Getenv("POSTGRES_HOST")
		if err != nil {
			log.Fatalf("Error loading .env files")
		}
	}

	dbUser, dbName, dbPassword :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD")

	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, 5432, dbUser, dbPassword, dbName)

	db, err := gorm.Open("postgres", postgresInfo)

	fmt.Println(err)
	if err != nil {
		log.Println("Database connection errror", err)
		panic(err)
	}

	log.Println("Database connection established")
	// defer db.Close()

	setupDB(db)

	return db
}
