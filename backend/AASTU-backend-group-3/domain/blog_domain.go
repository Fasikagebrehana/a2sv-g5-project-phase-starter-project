package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrBlogNotFound = errors.New("blog not found")
type Blog struct {
    ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`           // Unique identifier for the blog post
    Title         string    `json:"title"`          // Title of the blog post
    Content       string    `json:"content"`        // Content of the blog post
    AuthorID      string    `json:"author_id"`      // ID of the user who created the post
	AutorName     string    `json:"author_name"`    // Name of the user who created the post
    Tags          []string  `json:"tags"`           // Tags associated with the blog post
    CreatedAt     primitive.Timestamp `bson:"createdAt" json:"createdAt"`
	UpdatedAt     primitive.Timestamp `bson:"updatedAt" json:"updatedAt"`
    LikesCount    int       `json:"likes_count"`    // Number of likes the blog post has received
    DislikesCount int       `json:"dislikes_count"` // Number of dislikes the blog post has received
    ViewCount     int       `json:"view_count"`     // Number of views the blog post has received
    CommentsCount int       `json:"comments_count"` // Number of comments the blog post has received
	Visibility     string    `bson:"visibility"`
}



type BlogResponse struct {
	Title         string    `json:"title"`          // Title of the blog post
    Content       string    `json:"content"`        // Content of the blog post
    AuthorID      string    `json:"author_id"`      // ID of the user who created the post
    Tags          []string  `json:"tags"`           // Tags associated with the blog post
    CreatedAt 	  primitive.Timestamp `bson:"createdAt" json:"createdAt"`
	UpdatedAt 	  primitive.Timestamp `bson:"updatedAt" json:"updatedAt"`
    LikesCount    int       `json:"likes_count"`    // Number of likes the blog post has received
    DislikesCount int       `json:"dislikes_count"` // Number of dislikes the blog post has received
    ViewCount     int       `json:"view_count"`     // Number of views the blog post has received
    CommentsCount int       `json:"comments_count"` // Number of comments the blog post has received
	Visibility     string    `bson:"visibility"`
}




type BlogRepository interface {

	// Create and Get new blog post
	CreateBlog(username, userID string,blog Blog) (Blog, error)
	DeleteBlog( id string) (Blog, error)
	UpdateBlog( blog Blog, blogId string) (Blog, error)
	GetBlogByID( id string) (Blog, error)
	GetBlogs( page, limit int64, sortBy, tag, authorName string) ([]Blog, error)
	GetUserBlogs(userID string) ([]Blog, error)



	// Like, Dislike, View, Comment on a blog post
	// LikeBlog(userID, blogID string) error
	// DislikeBlog(userID, blogID string) error


	// // Retrieve a blog post by its ID
	// GetBlogByID(ctx context.Context, id string) (*Blog, error)
	
	// // Update an existing blog post
	// UpdateBlog(ctx context.Context, blog Blog) error
	
	// // Delete a blog post by its ID
	// DeleteBlog(ctx context.Context, id string) error
	
	// // Retrieve blog posts with pagination and sorting
	// GetBlogs(ctx context.Context, offset int64, limit int64, sortBy string) ([]Blog, error)
	
	// // Search for blog posts based on a query and additional filters
	// SearchBlogs(ctx context.Context, query string, filters map[string]interface{}) ([]Blog, error)
	
	// // Filter blog posts by tags, date, popularity, or other criteria
	// FilterBlogs(ctx context.Context, filters map[string]interface{}, sortBy string) ([]Blog, error)
	
	// // Track the popularity of a blog post (views, likes, dislikes, comments)
	// TrackPopularity(ctx context.Context, blogID string, action string) error
}

type BlogUsecase interface {
    CreateBlog(username, userID string, blog Blog) (Blog, error)    
	DeleteBlog(role, userId,id string) (Blog, error)
	UpdateBlog( blog Blog,role, blogId string)  (Blog, error)   
    GetBlogByID(id string) (Blog, error)     
	GetBlogs( page, limit int64, sortBy, tag, authorName string) ([]Blog, error)
	GetUserBlogs(userID string) ([]Blog, error)
	
	// Like, Dislike, View, Comment on a blog post
	// LikeBlog(userID, blogID string) error
	// DislikeBlog(userID, blogID string) error
	
	
	// Handle retrieval of a blog by ID
    // UpdateBlog(ctx context.Context, blog Blog) error                         // Handle updating a blog
    // DeleteBlog(ctx context.Context, id string) error                         // Handle deleting a blog
    // GetBlogs(ctx context.Context, offset int64, limit int64, sortBy string) ([]Blog, error) // Handle retrieval of multiple blogs with pagination and sorting
    // SearchBlogs(ctx context.Context, query string, filters map[string]interface{}) ([]Blog, error) // Handle search functionality
    // FilterBlogs(ctx context.Context, filters map[string]interface{}, sortBy string) ([]Blog, error) // Handle filtering of blog posts
    // TrackPopularity(ctx context.Context, blogID string, action string) error  // Handle tracking of blog post popularity
}




