package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog represents a blog post with flexible content.
type Blog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	Author    primitive.ObjectID `json:"ownerID" bson:"ownerID"`       // ID of the blog author
	Title     string             `json:"title" bson:"title"`           // Title of the blog
	Content   []interface{}      `json:"content" bson:"content"`       // Array of any type of content
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the blog was created
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"` // Timestamp for when the blog was last updated
	Tags      []BlogTag          `json:"tags" bson:"tags"`             // Tags for categorizing the blog
	Views     int64              `json:"views" bson:"views"`           // Count of views
	Likes     int64              `json:"likes" bson:"likes"`           // Count of likes
}

type BlogTag struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`    // Unique identifier for the blog
	Name string             `json:"name" bson:"name"` // Name of the blog
}

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of commenter
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`       // ID of blog
	Content   string             `json:"content" bson:"content"`       // Content of the comment
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the comment was created
	ReplyToId primitive.ObjectID `json:"reply_to" bson:"reply_to"`     // ID of comment

}

type Like struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of user
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`       // ID of blog
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the like was created
}

type View struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // Unique identifier for the blog
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of user
	BlogID    primitive.ObjectID `json:"blog_id" bson:"blog_id"`       // ID of blog
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"` // Timestamp for when the view was created
}

type BlogFilter struct {
	AuthorID  *primitive.ObjectID // Filter by Author ID
	Tags      []string            // Filter by Tags
	Title     *string             // Filter by Title (exact or partial match)
	DateRange *DateRange          // Filter by Creation Date Range
	Content   *string             // Filter by Content (exact or partial match)
	Keyword   *string             // Filter by keyword in title, content, or tags

}

// DateRange represents a time range for filtering
type DateRange struct {
	From time.Time // Start date for the range
	To   time.Time // End date for the range
}
