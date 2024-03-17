package infrastructure

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"module_blog.xws.com/model"
)

func InitDb() *gorm.DB {
	connectionString := formConnectionString()
	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil
	}

	database.AutoMigrate(&model.Blog{})
	return database
}

func formConnectionString() string {
	host := "localhost"
	port := 5432
	username := "postgres"
	password := "super"
	dbname := "explorer-v1"

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, password, host, port, dbname)

	return connectionString
}
