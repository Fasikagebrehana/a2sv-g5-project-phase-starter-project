package dto

import (
	"time"

	"aait.backend.g10/domain"
	"github.com/google/uuid"
)

// A dto for blog where it gives the blog details and the author details

type BlogDto struct {
	ID         		uuid.UUID   `json:"id"`
	Title      		string      `json:"title"`
	Content    		string      `json:"content"`
	Author     		uuid.UUID   `json:"author"`
	Tags       		[]string    `json:"tags"`
	CreatedAt  		time.Time   `json:"createdAt"`		
	UpdatedAt  		time.Time   `json:"updatedAt"`
	ViewCount  		int         `json:"viewCount"`
	LikeCount  		int		    `json:"likeCount"`
	DislikeCount	int		    `json:"dislikeCount"`
	CommentCount	int	   		`json:"commentCount"`
	AuthorName 		string      `json:"authorName"`
}


func NewBlogDto(blog domain.Blog, author domain.User, likeCount int, disLikeCount int, commentCount int) *BlogDto {
	return &BlogDto{
		ID:         blog.ID,
		Title:      blog.Title,
		Content:    blog.Content,
		Author:     blog.Author,
		Tags:       blog.Tags,
		CreatedAt:  blog.CreatedAt,
		UpdatedAt:  blog.UpdatedAt,
		ViewCount:  blog.ViewCount,
		LikeCount:  likeCount,
		DislikeCount: disLikeCount,
		CommentCount: commentCount,
		AuthorName: author.FullName,
	}
}