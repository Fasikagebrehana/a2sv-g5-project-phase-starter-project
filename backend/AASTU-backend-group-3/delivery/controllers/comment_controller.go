package controllers

import (
	"net/http"

	"group3-blogApi/domain"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentController struct {
	commentUsecase domain.CommentUsecase
}

func NewCommentController(commentUsecase domain.CommentUsecase) *CommentController {
	return &CommentController{
		commentUsecase: commentUsecase,
	}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var comment domain.Comment
	comment.UserID = userObjID
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	createdComment, err := c.commentUsecase.CreateComment(&comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Comment created successfully",
		
		"data": createdComment,
	})
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {

	var comment domain.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentID := ctx.Param("id")
	userID := ctx.GetString("user_id")
	Roles := ctx.GetString("role")
	
	commentIDObj, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.ID = commentIDObj

	updatedComment, err := c.commentUsecase.UpdateComment(&comment, Roles, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comment updated successfully",
		
		"data": updatedComment,
	})
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("id")
	userID := ctx.GetString("user_id")
	Roles := ctx.GetString("role")
	

	deletedComment, err := c.commentUsecase.DeleteComment(commentID, Roles, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully", 
		"data": deletedComment,
	})
}

func (c *CommentController) GetCommentByID(ctx *gin.Context) {
	commentID := ctx.Param("id")

	comment, err := c.commentUsecase.GetCommentByID(commentID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	postID := ctx.Param("postID")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")


	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	comments, err := c.commentUsecase.GetComments(postID, pageInt, limitInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c *CommentController) CreateReply(ctx *gin.Context) {
	var reply domain.Reply
	if err := ctx.ShouldBindJSON(&reply); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdReply, err := c.commentUsecase.CreateReply(&reply)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Reply created successfully",
		"data": createdReply,
	})
}

func (c *CommentController) UpdateReply(ctx *gin.Context) {
	var reply domain.Reply
	if err := ctx.ShouldBindJSON(&reply); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	replyID := ctx.Param("id")
	userID := ctx.GetString("user_id")

	replyIDObj, err := primitive.ObjectIDFromHex(replyID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	reply.ID = replyIDObj

	updatedReply, err := c.commentUsecase.UpdateReply(&reply, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Reply updated successfully",
		"data": updatedReply,
	})
}

func (c *CommentController) DeleteReply(ctx *gin.Context) {
	replyID := ctx.Param("id")
	userID := ctx.GetString("user_id")
	Roles := ctx.GetString("role")

	deletedReply, err := c.commentUsecase.DeleteReply(replyID, Roles, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Reply deleted successfully",
		"data": deletedReply,
	})
}

func (c *CommentController) GetReplies(ctx *gin.Context) {
	commentID := ctx.Param("commentID")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	replies, err := c.commentUsecase.GetReplies(commentID, pageInt, limitInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Replies retrieved successfully",
		"data": replies,
	})
}

func (c *CommentController) LikeComment(ctx *gin.Context) {
	commentID := ctx.Param("id")
	userID := ctx.GetString("user_id")


	if err := c.commentUsecase.LikeComment(commentID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *CommentController) UnlikeComment(ctx *gin.Context) {
	commentID := ctx.Param("id")
	userID := ctx.GetString("user_id")

	if err := c.commentUsecase.UnlikeComment(commentID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *CommentController) LikeReply(ctx *gin.Context) {
	replyID := ctx.Param("id")
	userID := ctx.GetString("user_id")

	if err := c.commentUsecase.LikeReply(replyID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *CommentController) UnlikeReply(ctx *gin.Context) {
	replyID := ctx.Param("id")
	userID := ctx.GetString("user_id")

	if err := c.commentUsecase.UnlikeReply(replyID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
