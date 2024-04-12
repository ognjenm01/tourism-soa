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
	//dburi := os.Getenv("MONGO_DB_URI")
	dburi := "mongodb://mongo:27017/"

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
	/*host := os.Getenv("BLOGS_DB_HOST")
	port := os.Getenv("BLOGS_DB_PORT")
	username := os.Getenv("BLOGS_DB_USERNAME")
	password := os.Getenv("BLOGS_DB_PASSWORD")
	dbname := os.Getenv("BLOGS_DB")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname)*/
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

	port := os.Getenv("BLOGS_DB_PORT")
	if port == "" {
		port = "5432"
	}

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, port)

	return connectionString
}

func clearDB(database *gorm.DB) {
	database.Exec("TRUNCATE blog.blogs RESTART IDENTITY CASCADE;")
}
