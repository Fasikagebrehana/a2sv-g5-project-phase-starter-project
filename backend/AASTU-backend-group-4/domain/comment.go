package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BlogID    primitive.ObjectID `bson:"blog_id" json:"blog_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
type CommentRepository interface {
	GetCommentsCount(ctx context.Context, blogID primitive.ObjectID) (int, error)
	CreateComment(ctx context.Context, comment Comment) error
	GetCommentByID(ctx context.Context, id primitive.ObjectID) (Comment, error)
	GetCommentsByBlogID(ctx context.Context, blogID primitive.ObjectID) ([]Comment, error)
	DeleteComment(ctx context.Context, commentID primitive.ObjectID) error
}

type CommentRequest struct {
	BlogID  primitive.ObjectID `bson:"blog_id" json:"blog_id" binding:"required"`
	UserID  primitive.ObjectID `bson:"user_id" json:"user_id"`
	Content string             `bson:"content" json:"content" binding:"required"`
}
