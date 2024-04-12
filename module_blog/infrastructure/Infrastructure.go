package infrastructure

import (
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"module_blog.xws.com/model"
)

func InitDb() (*mongo.Client, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func InitDb1() *gorm.DB {
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

	clearDB(database)

	return database
}

func formConnectionString() string {
	host := os.Getenv("BLOGS_DB_HOST")
	port := os.Getenv("BLOGS_DB_PORT")
	username := os.Getenv("BLOGS_DB_USERNAME")
	password := os.Getenv("BLOGS_DB_PASSWORD")
	dbname := os.Getenv("BLOGS_DB")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname)

	return connectionString
}

func clearDB(database *gorm.DB) {
	database.Exec("TRUNCATE blog.blogs RESTART IDENTITY CASCADE;")
}
