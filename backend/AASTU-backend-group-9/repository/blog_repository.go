package repository

import (
	"blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"blog/database"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type blogRepository struct {
	database   database.Database
	collection string
}

// NewBlogRepository returns a new instance of blogRepository implementing the domain.BlogRepository interface.
func NewBlogRepository(db database.Database, collection string) domain.BlogRepository {
	return &blogRepository{
		database:   db,
		collection: collection,
	}
}

// CreateBlog inserts a new blog into the MongoDB collection.
func (r *blogRepository) CreateBlog(ctx context.Context, blog *domain.Blog) error {
	collection := r.database.Collection(r.collection)
	_, err := collection.InsertOne(ctx, blog)
	return err
}

// GetBlogByID fetches a blog by its ID from the MongoDB collection.
func (r *blogRepository) GetBlogByID(ctx context.Context, id primitive.ObjectID) (*domain.Blog, error) {
	var blog domain.Blog
	collection := r.database.Collection(r.collection)

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&blog)
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

// GetAllBlogs fetches all blogs with pagination and sorting.
func (r *blogRepository) GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*domain.Blog, error) {
	var blogs []*domain.Blog

	// Pagination logic
	skip := (page - 1) * limit
	findOptions := options.Find().SetSort(bson.D{{Key: sortBy, Value: -1}}).SetSkip(int64(skip)).SetLimit(int64(limit))

	collection := r.database.Collection(r.collection)
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	if err := cursor.Close(ctx); err != nil {
		return nil, err
	}

	return blogs, nil
}

// UpdateBlog updates a blog in the MongoDB collection.
func (r *blogRepository) UpdateBlog(ctx context.Context, blog *domain.Blog) error {
	collection := r.database.Collection(r.collection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": blog.ID}, bson.M{"$set": blog})
	return err
}

// DeleteBlog deletes a blog by its ID from the MongoDB collection.
func (r *blogRepository) DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	collection := r.database.Collection(r.collection)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// Repositories/blog_repository.go
// SearchBlogs searches for blogs based on query and filters.
func (r *blogRepository) SearchBlogs(ctx context.Context, title string, author string) (*[]domain.Blog, error) {
	filter := bson.M{}

	// Add search filters based on the provided title and author
	if title != "" {
		filter["title"] = bson.M{"$regex": title, "$options": "i"}
	}
	if author != "" {
		filter["author"] = bson.M{"$regex": author, "$options": "i"}
	}

	var blogs []domain.Blog
	collection := r.database.Collection(r.collection)
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode the cursor results into the blogs slice
	if err = cursor.All(context.Background(), &blogs); err != nil {
		return nil, err
	}

	return &blogs, nil

}


func (r *blogRepository) FilterBlogs(ctx context.Context, popularity string, tags []string, startDate string, endDate string) ([]*domain.Blog, error) {
	var blogs []*domain.Blog

	// Define a filter without conditions initially
	filter := bson.M{}

	if len(tags) > 0 {
		filter["tags"] = bson.M{"$in": tags}
	}

	// Add filters based on date (assuming date is stored as a string, adjust if necessary)
	if startDate != "" && endDate != "" {
		filter["date"] = bson.M{
			"$gte": startDate,
			"$lte": endDate,
		}
	} else if startDate != "" {
		filter["date"] = bson.M{
			"$gte": startDate,
		}
	} else if endDate != "" {
		filter["date"] = bson.M{
			"$lte": endDate,
		}
	}

	// Define sort options based on popularity
	sortOptions := bson.D{}

	if popularity != "" {
		switch popularity {
		case "most_viewed":
			// Sort by views in descending order
			sortOptions = bson.D{{Key: "views", Value: -1}}
		case "most_liked":
			// Sort by likes in descending order
			sortOptions = bson.D{{Key: "likes", Value: -1}}
		}
	}

	collection := r.database.Collection(r.collection)
	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(sortOptions))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	return blogs, nil
}


func (r *blogRepository) AddComment(ctx context.Context, id primitive.ObjectID, comment *domain.Comment) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"comments": comment}}
	_, err := r.database.Collection(r.collection).UpdateOne(ctx, filter, update)
	return err
}



func (r *blogRepository) HasUserDisliked(ctx context.Context, id primitive.ObjectID, userID string) (bool, error) {
	filter := bson.M{"_id": id, "dislikes": userID}
	count, err := r.database.Collection(r.collection).CountDocuments(ctx, filter)
	return count > 0, err
}



func (r *blogRepository) IncrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{metric: 1}}
	_, err := r.database.Collection(r.collection).UpdateOne(ctx, filter, update)
	return err
}

func (r *blogRepository) DecrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{metric: -1}}
	_, err := r.database.Collection(r.collection).UpdateOne(ctx, filter, update)
	return err
}
