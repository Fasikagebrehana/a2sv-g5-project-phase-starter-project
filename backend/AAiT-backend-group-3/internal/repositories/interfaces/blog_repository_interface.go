package repository_interface

import (
	"AAIT-backend-group-3/internal/domain/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRepositoryInterface interface {
	CreateBlog(blog *models.Blog) error
	GetBlogByID(blogID primitive.ObjectID) (*models.Blog, error)
	GetBlogs(filter map[string]interface{}, search string, page int, limit int) ([]*models.Blog, error)
	EditBlog(logID primitive.ObjectID, newBlog *models.Blog) error
	DeleteBlog(blogID primitive.ObjectID) error
	AddCommentToTheList(blogID primitive.ObjectID, commentID primitive.ObjectID) error
}
