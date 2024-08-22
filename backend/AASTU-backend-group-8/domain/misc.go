package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	// ID        uint      `json:"id" gorm:"primaryKey"`
	ID   primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name string             `json:"name" validate:"required"`
}

type Pagination struct {
	Page  int `json:"page" validate:"min=1"`
	Limit int `json:"limit" validate:"min=1,max=100"`
}

type BlogFilter struct {
	Tags      []string `json:"tags"`
	AuthorID  string   `json:"author_id"`
	Search    string   `json:"search"`
	Title     string   `json:"title"`
	DateRange struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	} `json:"date_range"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	BlogID    primitive.ObjectID `bson:"blog_id"`
	UserID    primitive.ObjectID `bson:"author_id"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
}
type Like struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	BlogID    primitive.ObjectID `bson:"blog_id"`
	UserID    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
}
type LikeUsecaseInterface interface {
	AddLike(blogID, userID primitive.ObjectID) error
	RemoveLike(likeID primitive.ObjectID) error
	GetLikesByBlogID(blogID primitive.ObjectID) ([]Like, error)
}
type CommentUsecaseInterface interface {
	AddComment(comment *Comment) error
	GetCommentsByBlogID(blogID primitive.ObjectID) ([]Comment, error)
	UpdateComment(commentID primitive.ObjectID, content string) error
	DeleteComment(commentID primitive.ObjectID) error
}

// // IsExpired checks if the refresh token is expired
// func (r *RefreshToken) IsExpired() bool {
// 	return time.Now().After(r.ExpiresAt)
// }

//todo:
// the AI model
// the token models
