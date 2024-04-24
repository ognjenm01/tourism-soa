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

type BlogCommentRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *BlogCommentRepository) GetById(id string) (model.BlogComment, error) {
	var blogComment model.BlogComment

	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogComments")

	// Define a filter for finding the blog by id
	filter := bson.M{"_id": id}

	// Perform the find operation
	err := collection.FindOne(context.Background(), filter).Decode(&blogComment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return blogComment, nil
		}
		return blogComment, err
	}

	return blogComment, nil

	/*blogComment := model.BlogComment{}
	pk, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return blogComment, err
	}

	result := repo.DatabaseConnection.First(&blogComment, "id = ?", pk)

	if result != nil {
		return blogComment, result.Error
	}

	return blogComment, nil*/
}

func (repo *BlogCommentRepository) GetAll() ([]model.BlogComment, error) {
	var blogComments []model.BlogComment

	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogComments")

	// Perform the find operation
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document into a blog object
	for cursor.Next(context.Background()) {
		var blogComment model.BlogComment
		if err := cursor.Decode(&blogComment); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		blogComments = append(blogComments, blogComment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogComments, nil

	/*blogComments := []model.BlogComment{}
	result := repo.DatabaseConnection.Find(&blogComments)

	if result != nil {
		return blogComments, result.Error
	}

	return blogComments, nil*/
}

func (repo *BlogCommentRepository) Create(blogComment *model.BlogComment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogComments")

	// Perform the insert operation
	_, err := collection.InsertOne(ctx, &blogComment)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil

	/*result := repo.DatabaseConnection.Create(blogComment)

	if result.Error != nil {
		return result.Error
	}

	return nil*/
}

func (repo *BlogCommentRepository) Update(blogComment *model.BlogComment) error {
	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogComments")

	// Define the filter for finding the blogComment by its ID
	filter := bson.M{"_id": blogComment.Id}

	// Define the update operation
	update := bson.M{
		"$set": bson.M{
			"Comment": blogComment.Comment,
			//"SystemStatus": blog.SystemStatus,
		},
	}

	// Perform the update operation
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil

	/*result := repo.DatabaseConnection.Save(blogComment)

	if result.Error != nil {
		return result.Error
	}

	return nil*/
}

func (repo *BlogCommentRepository) Delete(id string) error {
	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogComments")

	// Define the filter for finding the blogComment by its ID
	filter := bson.M{"_id": id}

	// Perform the delete operation
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil

	/*result := repo.DatabaseConnection.Delete(model.BlogComment{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil*/
}

func (repo *BlogCommentRepository) GetCommentsByBlogId(id string) ([]model.BlogComment, error) {
	blogComments := []model.BlogComment{}

	collection := repo.DatabaseConnection.Database("mongoDemo").Collection("blogComments")

	// Define a filter for finding the blog comments by blog id
	filter := bson.M{"blog_id": id}

	// Perform the find operation
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return blogComments, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var blogComment model.BlogComment
		err := cur.Decode(&blogComment)
		if err != nil {
			return blogComments, err
		}
		blogComments = append(blogComments, blogComment)
	}

	if err := cur.Err(); err != nil {
		return blogComments, err
	}

	return blogComments, nil

	/*blogComments := []model.BlogComment{}
	result := repo.DatabaseConnection.Raw("SELECT * FROM blog.blog_comments WHERE blog.blog_comments.blog_id = ?", id).Scan(&blogComments)
	if result.Error != nil {
		return blogComments, result.Error
	}
	return blogComments, result.Error*/
}
