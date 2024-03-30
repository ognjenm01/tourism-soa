package test_infrastructure

import (
	"fmt"
	"tour/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func createConnectionString() string {
	host := "localhost"
	username := "postgres"
	password := "super"
	dbname := "tour_test_db"
	port := "5400"

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

	database.AutoMigrate(&model.Tour{})
	database.AutoMigrate(&model.TourTag{})
	database.AutoMigrate(&model.Keypoint{})
	database.AutoMigrate(&model.TourReview{})
	database.AutoMigrate(&model.ImageLink{})
	database.AutoMigrate(&model.TourEquipment{})
	database.AutoMigrate(&model.Equipment{})

	clearDB(database)
	fillDB(database)
	return database
}

func clearDB(database *gorm.DB) {
	database.Exec("TRUNCATE tours.tours RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.tour_reviews RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.keypoints RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.tour_tags RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.image_links RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.tour_equipments RESTART IDENTITY CASCADE;")
	database.Exec("TRUNCATE tours.equipment RESTART IDENTITY CASCADE;")
}

func fillDB(database *gorm.DB) {
	database.Exec("INSERT INTO tours.tours(user_id, name, description, price, difficulty, transport_type, status, duration, distance, status_update_time) VALUES ('1', 'Fantastic tour to Serbia', 'Exceptional', '300', '0', '0', '0', '300', '300', '2004-10-19 10:23:54+02');")
	database.Exec("INSERT INTO tours.tours(user_id, name, description, price, difficulty, transport_type, status, duration, distance, status_update_time) VALUES ('-1', 'Tour to Spain', 'Fenomenal', '400', '1', '0', '1', '300', '300', '2004-10-20 11:23:54+02');")
	database.Exec("INSERT INTO tours.tours(user_id, name, description, price, difficulty, transport_type, status, duration, distance, status_update_time) VALUES ('-1', 'Tour to Argentina', 'Wonderful', '500', '1', '2', '0', '300', '300', '2004-10-21 12:23:54+02');")

	database.Exec("INSERT INTO tours.tour_tags(tour_id, tag) VALUES ('1', 'Hot');")

	database.Exec("INSERT INTO tours.keypoints(tour_id, name, latitude, longitude, description, position, image, secret) VALUES (1, 'Nova tačka 1', 54.701251, 9.876543, 'Opis nove tačke 1', 1, 'new_image_1.jpg', 'new_secret_1');")
	database.Exec("INSERT INTO tours.keypoints(tour_id, name, latitude, longitude, description, position, image, secret) VALUES (2, 'Nova tačka 2', 55.812361, 8.765432, 'Opis nove tačke 2', 1, 'new_image_2.jpg', 'new_secret_2');")
	database.Exec("INSERT INTO tours.keypoints(tour_id, name, latitude, longitude, description, position, image, secret) VALUES (2, 'Nova tačka 3', 56.923471, 7.654321, 'Opis nove tačke 3', 2, 'new_image_3.jpg', 'new_secret_3');")

	database.Exec("INSERT INTO tours.equipment(name, description) VALUES ('TestEq1', 'TestEqDesc1');")
	database.Exec("INSERT INTO tours.equipment(name, description) VALUES ('TestEq2', 'TestEqDesc2');")
	database.Exec("INSERT INTO tours.equipment(name, description) VALUES ('TestEq3', 'TestEqDesc3');")

	database.Exec("INSERT INTO tours.tour_equipments(tour_id, equipment_id) VALUES (1, 1);")
	database.Exec("INSERT INTO tours.tour_equipments(tour_id, equipment_id) VALUES (1, 2);")
	database.Exec("INSERT INTO tours.tour_equipments(tour_id, equipment_id) VALUES (2, 1);")
	database.Exec("INSERT INTO tours.tour_equipments(tour_id, equipment_id) VALUES (3, 1);")
}
