package repo

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"module_blog.xws.com/model"
)

type BlogRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *BlogRepository) GetByPeopleUFollow() ([]model.Blog, error) {
	var blogs []model.Blog

	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogs")
	ctx := context.Background()

	// Find all documents in the collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error finding documents:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog model.Blog
		if err := cursor.Decode(&blog); err != nil {
			fmt.Println("Error decoding document:", err)
			continue
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (repo *BlogRepository) GetById(id string) (model.Blog, error) {
	var blog model.Blog

	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogs")

	// Define a filter for finding the blog by id
	kita, err2 := strconv.Atoi(id)
	if err2 != nil {
		log.Fatalln("prema meni bezdusaaaaaaan")
	}
	filter := bson.M{"kita": kita}

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

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

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

func (repo *BlogRepository) Update(blog *model.Blog) (*model.Blog, error) {
	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogs") // Specify your database and collection name

	filter := bson.M{"_id": blog.Kita}

	// Define the update operation
	update := bson.M{
		"$set": bson.M{
			"Description":  blog.Description,
			"SystemStatus": blog.SystemStatus,
		},
	}

	// Perform the update operation
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := collection.FindOneAndUpdate(context.Background(), filter, update, options)

	if result.Err() != nil {
		log.Fatalln(result.Err().Error())
		return nil, result.Err()
	}

	var updatedBlog model.Blog

	if err := result.Decode(&updatedBlog); err != nil {
		return nil, err
	}

	return &updatedBlog, nil
}

func (repo *BlogRepository) Delete(id string) error {
	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogs") // Specify your database and collection name

	kita, err2 := strconv.Atoi(id)
	if err2 != nil {
		log.Fatalln("prema meni bezdusaaaaaaan")
	}
	filter := bson.M{"kita": kita}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
