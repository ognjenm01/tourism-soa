package repo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	//"gorm.io/gorm"
	"module_blog.xws.com/model"
)

type BlogRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *BlogRepository) GetById(id string) (model.Blog, error) {
	var blog model.Blog

	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogs")

	// Define a filter for finding the blog by id
	filter := bson.M{"_id": id}

	// Perform the find operation
	err := collection.FindOne(context.Background(), filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return blog, nil // Return empty blog if not found
		}
		return blog, err
	}

	return blog, nil
}

func (repo *BlogRepository) GetAll() ([]model.Blog, error) {
	var blogs []model.Blog

	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogs") // Specify your database and collection name

	// Perform the find operation
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document into a blog object
	for cursor.Next(context.Background()) {
		var blog model.Blog
		if err := cursor.Decode(&blog); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (repo *BlogRepository) Create(blog *model.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogs")

	// Perform the insert operation
	_, err := collection.InsertOne(ctx, &blog)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (repo *BlogRepository) Update(blog *model.Blog) error {
	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogs") // Specify your database and collection name

	// Define the filter for finding the blog by its ID
	filter := bson.M{"_id": blog.Id}

	// Define the update operation
	update := bson.M{
		"$set": bson.M{
			"Description":  blog.Description,
			"SystemStatus": blog.SystemStatus,
		},
	}

	// Perform the update operation
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *BlogRepository) Delete(id string) error {
	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogs") // Specify your database and collection name

	// Define the filter for finding the blog by its ID
	filter := bson.M{"_id": id}

	// Perform the delete operation
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
