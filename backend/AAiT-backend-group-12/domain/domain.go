package domain

import (
	"blog_api/domain/dtos"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

/*
Defines the names of the collections in the DB
*/
const (
	CollectionUsers = "users"
	CollectionBlogs = "blogs"
)

const (
	RoleUser = "user" 
	RoleAdmin = "admin"
	RoleRoot = "root"
)


const (
	VerifyEmailType   = "verify_email"
	ResetPasswordType = "reset_password"
)

type Response gin.H

type VerificationData struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	Type      string    `json:"type"`
}


type User struct {
	Username         string           `json:"username"`
	Email            string           `json:"email"`
	Password         string           `json:"password"`
	PhoneNumber      string           `json:"phone_number"`
	Bio              string           `json:"bio"`
	Role             string           `json:"role"`
	CreatedAt        time.Time        `json:"created_at"`
	RefreshToken     string           `json:"refresh_token"`
	IsVerified       bool             `json:"is_verified"`
	VerificationData VerificationData `json:"verification_data"`
}


type Blog struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Username   string    `json:"username"`
	Tags       []string  `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ViewCount  uint      `json:"view_count"`
	LikedBy    []string  `json:"liked_by"`
	DislikedBy []string  `json:"disliked_by"`
	Comments   []Comment `json:"comment"`
}


type NewBlog struct{
	Title      string    `json:"title" validate:"required,MinWord=1"`
	Content    string    `json:"content" validate:"required,MinWord=25"`
	Tags       []string  `json:"tags"`
}


type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type NewComment struct{
	Content   string    `json:"content" validate:"required,min=3"`
}


type BlogFilterOptions struct {
	Title         string // Search by title
	Author        string // Search by author name
	Tags          []string
	DateFrom      time.Time
	DateTo        time.Time
	SortBy        string // Sort by criteria: date, like count, dislike count, view count
	SortDirection string // Sort direction: asc, desc
	Page          int    // Pagination: Page number
	PostsPerPage  int    // Pagination: Posts per page
	MinLikes      int    // Filter by minimum likes
	MinDislikes   int    // Filter by minimum dislikes
	MinComments   int    // Filter by minimum comments
	MinViewCount  int    // Filter by minimum view count
}


type BlogRepositoryInterface interface {
	//Blog related methods
	FetchBlogPostByID(ctx context.Context, postID string) (*Blog, CodedError)
	FetchBlogPosts(ctx context.Context, filters BlogFilterOptions) ([]Blog, int, CodedError)
	InsertBlogPost(ctx context.Context, blog *Blog) CodedError
	UpdateBlogPost(ctx context.Context, id string, blog *NewBlog) CodedError
	DeleteBlogPost(ctx context.Context, id string) CodedError
	TrackBlogPopularity(ctx context.Context, blogId string, action string, username string) CodedError

	//Comment related methods
	// FetchComment(ctx context.Context, commentID, blogID string) (Comment, CodedError)
	CreateComment(ctx context.Context, comment *Comment, blogID, createdBy string) CodedError
	UpdateComment(ctx context.Context, comment *NewComment, commentID, blogID, userName string) CodedError
	DeleteComment(ctx context.Context, commentID, blogID, userName string) CodedError
}


type BlogUseCaseInterface interface {
	//Blog related methods
	GetBlogPostByID(ctx context.Context, id string) (*Blog, CodedError)
	GetBlogPosts(ctx context.Context, filters BlogFilterOptions) ([]Blog, int, CodedError)
	CreateBlogPost(ctx context.Context, blog *NewBlog, createdBy string) CodedError
	EditBlogPost(ctx context.Context, id string, blog *NewBlog, editedBy string) CodedError
	DeleteBlogPost(ctx context.Context, id, deletedBy string) CodedError
	TrackBlogPopularity(ctx context.Context, blogId, action, username string) CodedError
	GenerateTrendingTopics(keywords []string) ([]string, error)
	ReviewBlogContent(blogContent string) (string, error)
	GenerateBlogContent(topics []string) (string, error)

	//Comment related methods
	// FindComment(ctx context.Context, commentID, blogID string) (Comment, CodedError)
	AddComment(ctx context.Context, blogID string, newComment *NewComment, username string) CodedError
	UpdateComment(ctx context.Context, blogID string, commentID string, comment *NewComment, username string) CodedError
	DeleteComment(ctx context.Context, blogID, commentID, username string) CodedError
}


type UserRepositoryInterface interface {
	CreateUser(c context.Context, user *User) CodedError
	FindUser(c context.Context, user *User) (User, CodedError)
	SetRefreshToken(c context.Context, user *User, newRefreshToken string) CodedError
	UpdateUser(c context.Context, username string, user *dtos.UpdateUser) (map[string]string, CodedError)
	ChangeRole(c context.Context, username string, newRole string) CodedError
	VerifyUser(c context.Context, username string) CodedError
	UpdateVerificationDetails(c context.Context, username string, verificationData VerificationData) CodedError
}


type UserUsecaseInterface interface {
	Signup(c context.Context, user *User, hostUrl string) CodedError
	Login(c context.Context, user *User) (string, string, CodedError)
	RenewAccessToken(c context.Context, refreshToken string) (string, CodedError)
	UpdateUser(c context.Context, requestUsername string, tokenUsername string, user *dtos.UpdateUser) (map[string]string, CodedError)
	PromoteUser(c context.Context, username string) CodedError
	DemoteUser(c context.Context, username string) CodedError
	VerifyEmail(c context.Context, username string, token string, hostUrl string) CodedError
}
