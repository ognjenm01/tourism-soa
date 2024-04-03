package infrastructure

import (
	"fmt"
	"os"
	"tour/model"

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

	host := os.Getenv("TOURS_DB_HOST")
	if host == "" {
		host = "localhost"
	}

	username := os.Getenv("TOURS_DB_USERNAME")
	if username == "" {
		username = "postgres"
	}

	password := os.Getenv("TOURS_DB_PASSWORD")
	if password == "" {
		password = "super"
	}

	dbname := os.Getenv("TOURS_DB")
	if dbname == "" {
		dbname = "explorer-v1"
	}

	port := os.Getenv("TOURS_DB_PORT")
	if port == "" {
		port = "5432"
	}

	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)
	return connectionStr
}

func createTablePrefix() string {
	prefix := "tours."
	if prefix == "" {
		prefix = "tours."
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
	database.Exec("TRUNCATE tours.tours RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.tour_reviews RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.keypoints RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.tour_tags RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.equipment RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.tour_equipments RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.tour_progresses RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.tourist_positions RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.image_links RESTART IDENTITY CASCADE;")
}

func fillDB(database *gorm.DB) {

	database.AutoMigrate(&model.Tour{})
	database.Exec("INSERT INTO tours.tours(user_id, name, description, price, difficulty, transport_type, status, duration, distance, status_update_time) VALUES ('2', 'Fantastic tour to Serbia', 'Exceptional', '300', '0', '0', '0', '300', '300', '2004-10-19 10:23:54+02');")
	database.Exec("INSERT INTO tours.tours(user_id, name, description, price, difficulty, transport_type, status, duration, distance, status_update_time) VALUES ('2', 'Tour to Spain', 'Fenomenal', '400', '1', '0', '1', '300', '300', '2004-10-20 11:23:54+02');")
	database.Exec("INSERT INTO tours.tours(user_id, name, description, price, difficulty, transport_type, status, duration, distance, status_update_time) VALUES ('2', 'Tour to Argentina', 'Wonderful', '500', '1', '2', '0', '300', '300', '2004-10-21 12:23:54+02');")

	database.AutoMigrate(&model.TourTag{})
	database.Exec("INSERT INTO tours.tour_tags(tour_id, tag) VALUES ('1', 'Hot');")

	database.AutoMigrate(&model.Keypoint{})
	database.AutoMigrate(&model.TourReview{})
	database.AutoMigrate(&model.ImageLink{})
	database.AutoMigrate(&model.Equipment{})
	database.AutoMigrate(&model.TourEquipment{})
	database.AutoMigrate(&model.TourProgress{})
	database.AutoMigrate(&model.TouristPosition{})
}
