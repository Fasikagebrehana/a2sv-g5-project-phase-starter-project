package controller

import (
	domain "AAiT-backend-group-8/Domain"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) GetComments(ctx *gin.Context) {
	blogID := ctx.Param("blogID")
	fmt.Print(blogID)
	comments, err := c.commentUseCase.GetComments(blogID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, comments)
}

// a function to get the token from the authorization header , and decode it to get the user id
func (c *Controller) CreateComment(ctx *gin.Context) {
	blogID := ctx.Param("blogID")
	token, err := c.ExtractToken(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, name, err := c.commentUseCase.DecodeToken(token)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var comment domain.Comment
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	primUserId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	comment.AuthorID = primUserId
	comment.AuthorName = name

	fmt.Println(comment)

	//!TODO add the name and id of the user to the comment
	if comment.Body == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "comment body is required"})
		return
	}
	err = c.commentUseCase.CreateComment(&comment, blogID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "comment created successfully"})
}

func (c *Controller) UpdateComment(ctx *gin.Context) {
	commentID := ctx.Param("commentID")
	var comment domain.Comment
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if comment.Body == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "comment body is required"})
		return
	}
	err := c.commentUseCase.UpdateComment(&comment, commentID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "comment updated successfully"})
}

func (c *Controller) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("commentID")
	err := c.commentUseCase.DeleteComment(commentID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "comment deleted successfully"})
}

func (c *Controller) ExtractToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", gin.Error{
			Err:  errors.New("authorization header not provided"),
			Type: gin.ErrorTypePublic,
		}
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", gin.Error{
			Err:  errors.New("invalid authorization header format"),
			Type: gin.ErrorTypePublic,
		}
	}
	return parts[1], nil
}
