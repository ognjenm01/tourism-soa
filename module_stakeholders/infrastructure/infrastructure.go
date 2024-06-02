package infrastructure

import (
	"fmt"
	"os"
	"stakeholder/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func createConnectionString() string {
	//TODO Remove this later, global env will exists
	/*err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}*/

	host := os.Getenv("STAKEHOLDERS_DB_HOST")
	if host == "" {
		host = "localhost"
	}

	username := os.Getenv("STAKEHOLDERS_DB_USERNAME")
	if username == "" {
		username = "postgres"
	}

	password := os.Getenv("STAKEHOLDERS_DB_PASSWORD")
	if password == "" {
		password = "super"
	}

	dbname := os.Getenv("STAKEHOLDERS_DB")
	if dbname == "" {
		dbname = "explorer-v1"
	}

	port := os.Getenv("STAKEHOLDERS_DB_PORT")
	if port == "" {
		port = "5432"
	}

	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)
	return connectionStr
}

func createTablePrefix() string {
	prefix := "stakeholders."
	if prefix == "" {
		prefix = "stakeholders."
	}
	return prefix
}

func InitDB() *gorm.DB {
	connectionStr := createConnectionString()
	database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   createTablePrefix(),
			SingularTable: false,
		}})
	if err != nil {
		print(err)
		return nil
	}

	clearDB(database)
	fillDB(database)
	return database
}

func clearDB(database *gorm.DB) {
	database.Exec("TRUNCATE stakeholders.users RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE stakeholders.people RESTART IDENTITY CASCADE;")
}

func fillDB(database *gorm.DB) {

	database.AutoMigrate(&model.Person{})

	database.AutoMigrate(&model.User{})
	//Todo nikad
	//database.Exec("INSERT INTO tours.tour_tags(tour_id, tag) VALUES ('1', 'Hot');")
}
