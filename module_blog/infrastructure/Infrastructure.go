package infrastructure

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"module_blog.xws.com/model"
)

func InitDb() *gorm.DB {
	connectionString := formConnectionString()
	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "blog.",
		},
	})

	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	database.AutoMigrate(&model.Blog{})
	database.AutoMigrate(&model.BlogRating{})
	database.AutoMigrate(&model.BlogStatus{})
	database.AutoMigrate(&model.BlogComment{})

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
