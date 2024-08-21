package repository

import (
	"context"
	"meleket/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRepository struct {
	collection domain.Collection
}

// // Search implements domain.BlogRepositoryInterface.
// func (r *BlogRepository) Search(title string) ([]domain.BlogPost, error) {
// 	panic("unimplemented")
// }

func NewBlogRepository(col domain.Collection) *BlogRepository {
	return &BlogRepository{collection: col}
}

func (r *BlogRepository) Save(blog *domain.BlogPost) (interface{},error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	createdBlog, err := r.collection.InsertOne(ctx, blog)
	return createdBlog.InsertedID,err
}

func (r *BlogRepository) GetAllBlog() ([]domain.BlogPost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var blogs []domain.BlogPost
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(context.TODO(), &blogs)

	return blogs, err
}

func (r *BlogRepository) GetBlogByID(id primitive.ObjectID) (*domain.BlogPost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var blog domain.BlogPost
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&blog)
	return &blog, err
}

func (r *BlogRepository) Update(blog *domain.BlogPost) (*domain.BlogPost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": blog.ID}, blog)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var decoded domain.BlogPost
	if err := result.Decode(&decoded); err != nil {
		return nil, err
	}

	blog.ID = decoded.ID
	blog.AuthorID = decoded.AuthorID

	return blog, nil
}

// Search function here

func (r *BlogRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
