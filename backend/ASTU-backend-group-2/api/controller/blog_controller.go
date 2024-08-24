package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/forms"
	"github.com/gin-gonic/gin"
)

// interface for blog controllers
type blogController interface {
	GetBlogs() gin.HandlerFunc
	GetBlog() gin.HandlerFunc
	CreateBlog() gin.HandlerFunc
	UpdateBlog() gin.HandlerFunc
	DeleteBlog() gin.HandlerFunc

	GetComments() gin.HandlerFunc
	CreateComment() gin.HandlerFunc
	GetComment() gin.HandlerFunc
	UpdateComment() gin.HandlerFunc
	DeleteComment() gin.HandlerFunc
	CreateLike() gin.HandlerFunc
}

type BlogController struct {
	BlogUsecase    entities.BlogUsecase
	CommentUsecase entities.CommentUsecase
	UserUsecase    entities.UserUsecase
	Env            *bootstrap.Env
}

func (bc *BlogController) GetBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {

		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		dateFrom, _ := time.Parse(time.RFC3339, c.Query("date_from"))
		dateTo, _ := time.Parse(time.RFC3339, c.Query("date_to"))
		tags := strings.Split(c.Query("tags"), ",")
		popularityFrom, _ := strconv.Atoi(c.Query("popularity_from"))
		popularityTo, _ := strconv.Atoi(c.Query("popularity_to"))
		//in go when u split an empty string
		//Because the returned array is not empty. First element of it is an empty string ""
		if len(tags) == 1 && tags[0] == "" {
			tags = []string{}
		}
		var blogFilter entities.BlogFilter
		log.Printf("%#v\n", tags)
		blogFilter = entities.BlogFilter{
			Title:          c.Query("title"),
			Search:         c.Query("search"),
			Tags:           tags,
			DateFrom:       dateFrom,
			DateTo:         dateTo,
			Limit:          limit,
			Pages:          page,
			PopularityFrom: popularityFrom,
			PopularityTo:   popularityTo,
		}

		blogs, pagination, err := bc.BlogUsecase.GetAllBlogs(context.Background(), blogFilter)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		res := entities.PaginatedResponse{
			Data:     blogs,
			MetaData: pagination,
		}

		c.JSON(200, res)
	}
}

func (bc *BlogController) GetBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		blog, err := bc.BlogUsecase.GetBlogByID(context.Background(), blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blog)
	}
}

func (bc *BlogController) CreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBlog forms.BlogForm
		if err := c.ShouldBindJSON(&newBlog); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		userID, exists := c.Get("x-user-id")

		if !exists {
			c.JSON(500, gin.H{"error": "User not found"})
			return
		}

		uid, ok := userID.(string)

		if !ok {
			c.JSON(500, gin.H{"error": "User not found"})
			return
		}

		fmt.Println("user id", uid)

		user, err := bc.UserUsecase.GetUserById(context.Background(), uid)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		blog, err := bc.BlogUsecase.CreateBlog(context.Background(), &newBlog, user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, blog)
	}
}

func (bc *BlogController) BatchCreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {

		var newBlogs []forms.BlogForm

		if err := c.ShouldBindJSON(&newBlogs); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		userID := c.MustGet("x-user-id").(string)

		user, err := bc.UserUsecase.GetUserById(context.Background(), userID)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		err = bc.BlogUsecase.BatchCreateBlog(context.Background(), &newBlogs, user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Blogs created successfully"})
	}
}

func (bc *BlogController) UpdateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		var updatedBlog forms.BlogForm
		if err := c.ShouldBindJSON(&updatedBlog); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		blog, err := bc.BlogUsecase.UpdateBlog(context.Background(), blogID, &updatedBlog)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blog)
	}
}

func (bc *BlogController) DeleteBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		err := bc.BlogUsecase.DeleteBlog(context.TODO(), blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, nil)
	}
}
func (bc *BlogController) GetByTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		tags, _ := c.GetQueryArray("tags")
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

		blogs, pagination, err := bc.BlogUsecase.GetByTags(context.TODO(), tags, limit, page)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"blogs": blogs, "pageination": pagination})
	}
}

func (bc *BlogController) GetbyPopularity() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		blogs, pagination, err := bc.BlogUsecase.GetByPopularity(context.Background(), limit, page)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"blogs": blogs, "pageination": pagination})
	}
}
func (bc *BlogController) SortByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		blogs, pagination, err := bc.BlogUsecase.SortByDate(context.Background(), limit, page)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"blogs": blogs, "pageination": pagination})
	}
}
func (bc *BlogController) GetComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

		comments, pageination, err := bc.CommentUsecase.GetComments(c, blogID, limit, page)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("[ctrl] blog id", blogID)

		c.JSON(http.StatusOK, gin.H{"comments": comments, "metadata": pageination})
	}
}
func (bc *BlogController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("x-user-id").(string)
		blogID := c.Param("id")
		var commentIn forms.CommentForm

		if err := c.BindJSON(&commentIn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("comment input:", bc.CommentUsecase)
		comment, err := bc.CommentUsecase.CreateComment(c, userID, blogID, &commentIn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"comment": comment})

	}
}
func (bc *BlogController) GetComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")

		comment, err := bc.CommentUsecase.GetComment(c, commentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"comment": comment})
	}
}
func (bc *BlogController) UpdateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")

		var commentUpd entities.CommentUpdate

		if err := c.BindJSON(&commentUpd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		comment, err := bc.CommentUsecase.UpdateComment(c.Request.Context(), commentID, &commentUpd)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"comment": comment})
	}
}

func (bc *BlogController) DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")

		err := bc.CommentUsecase.DeleteComment(c, commentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{})
	}
}
