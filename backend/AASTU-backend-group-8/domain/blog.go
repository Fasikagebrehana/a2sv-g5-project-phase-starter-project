package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	// ID          string      `json:"id" gorm:"primaryKey"`
	ID        primitive.ObjectID `json:"_id"  bson:"_id,omitempty"`
	Title     string             `json:"title" validate:"required,min=5,max=255"`
	Content   string             `json:"content" validate:"required"`
	AuthorID  primitive.ObjectID `json:"author_id"` // Foreign key to User model
	Tags      []string           `json:"tags" gorm:"type:text[]"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	//PublishedAt *time.Time `json:"published_at"`  Optional
}

type SearchBlogPost struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BlogUsecaseInterface interface {
	CreateBlogPost(blog *BlogPost) (interface{}, error)
	GetAllBlogPosts() ([]BlogPost, error)
	GetBlogByID(id primitive.ObjectID) (*BlogPost, error)
	UpdateBlogPost(id primitive.ObjectID, blog *BlogPost) (*BlogPost, error)
	// SearchBlogPosts(query *SearchBlogPost) ([]BlogPost, error) // Add this method
	DeleteBlogPost(id primitive.ObjectID) error
}

// domain/blog_repository_interface.go

type BlogRepositoryInterface interface {
	Save(blog *BlogPost) (interface{},error)
	GetAllBlog() ([]BlogPost, error)
	GetBlogByID(id primitive.ObjectID) (*BlogPost, error)
	Update(blog *BlogPost) (*BlogPost, error)
	// Search(title string) ([]BlogPost, error) // Add this method
	Delete(id primitive.ObjectID) error
}
