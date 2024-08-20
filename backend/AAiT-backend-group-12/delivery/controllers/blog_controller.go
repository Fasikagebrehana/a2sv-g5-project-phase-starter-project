package controllers

import (
	"blog_api/domain"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BlogController struct {
	blogUseCase domain.BlogUseCaseInterface
}

var validate = validator.New()

func NewBlogController(bu domain.BlogUseCaseInterface) *BlogController {
	return &BlogController{
		blogUseCase: bu,
	}
}

// CreateBlogHandler handles the HTTP request for creating a new blog post.
func (bc *BlogController) CreateBlogHandler(c *gin.Context) {
	var blog domain.NewBlog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": err.Error()})
		return
	}

	err := validate.Struct(blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": err.Error()})
		return
	}

	userName, exists := c.Keys["username"]
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"message": "coudn't find the username field"})
		return
	}
	created_By := userName.(string)
	newErr := bc.blogUseCase.CreateBlogPost(c, &blog, created_By)
	if newErr != nil {
		c.JSON(GetHTTPErrorCode(newErr), domain.Response{"error": newErr.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Response{"message": "blog created successfully"})
}

// UpdateBlogHandler handles the HTTP request to update a blog post.
func (bc *BlogController) UpdateBlogHandler(c *gin.Context) {
	blogId := c.Param("id")
	var blog domain.NewBlog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": err.Error()})
		return
	}
	userName, exists := c.Keys["username"]
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"message": "coudn't find the username field"})
		return
	}
	userNameStr := userName.(string)
	err := bc.blogUseCase.EditBlogPost(c, blogId, &blog, userNameStr)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, domain.Response{"message": "updated successfuly"})
}

// DeleteBlogHandler handles the HTTP DELETE request to delete a blog post.
func (bc *BlogController) DeleteBlogHandler(c *gin.Context) {
	blogId := c.Param("id")
	userName, exists := c.Keys["username"]
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"message": "coudn't find the username field"})
		return
	}
	userNameStr := userName.(string)

	err := bc.blogUseCase.DeleteBlogPost(c, blogId, userNameStr)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, domain.Response{"message": "deleted successfuly"})
}

// GetBlogHandler handles the HTTP GET request to retrieve a list of blog posts based on filters.
func (bc *BlogController) GetBlogHandler(c *gin.Context) {
	var filters domain.BlogFilterOptions
	if err := c.ShouldBindJSON(&filters); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Invalid query parameters"})
		return
	}

	blogs, total, err := bc.blogUseCase.GetBlogPosts(c, filters)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}
	
	if len(blogs) == 0 {
		c.JSON(404, domain.Response{"message": "No blog found"})
		return
	}

	currentPage := 1
	postsPerPage := 10
	if filters.Page != 0 {
		currentPage = filters.Page
	}
	if filters.PostsPerPage != 0 {
		postsPerPage = filters.PostsPerPage
	}

	c.JSON(http.StatusOK, gin.H{"total": total, "blogs": blogs, "currentPage": currentPage, "postsPerPage": postsPerPage})
}

// GetBlogByIDHandler handles the HTTP GET request to retrieve a single blog post by its ID.
func (bc *BlogController) GetBlogByIDHandler(c *gin.Context) {
	blogId := c.Param("id")
	blog, err := bc.blogUseCase.GetBlogPostByID(c, blogId)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// handles like request.
func (bc *BlogController) BlogLikeHandler(c *gin.Context) {
	var requestBody domain.LikeOrDislikeRequest
	
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Invalid input data"})
		return
	}
	
	// extract the username from the context
	userName, exists := c.Keys["username"]
	if !exists {
		c.JSON(http.StatusForbidden, domain.Response{"message": "User not found"})
		return
	}

	userNameStr := userName.(string)
	err := bc.blogUseCase.TrackBlogPopularity(c, requestBody.BlogID, "like",requestBody.State, userNameStr)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{"message": "Action applied successfully"})
}

// handles like request.
func (bc *BlogController) BlogDisLikeHandler(c *gin.Context) {
	var requestBody domain.LikeOrDislikeRequest
	
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{"error": "Invalid input data"})
		return
	}
	
	// extract the username from the context
	userName, exists := c.Keys["username"]
	if !exists {
		c.JSON(http.StatusForbidden, domain.Response{"message": "User not found"})
		return
	}

	userNameStr := userName.(string)
	err := bc.blogUseCase.TrackBlogPopularity(c, requestBody.BlogID, "dislike",requestBody.State, userNameStr)
	if err != nil {
		c.JSON(GetHTTPErrorCode(err), domain.Response{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{"message": "Action applied successfully"})
}

// GenerateContentHandler handles requests to generate content
func (bc *BlogController) GenerateContentHandler(c *gin.Context) {
	var req struct {
		Topics []string `json:"topics"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	content, err := bc.blogUseCase.GenerateBlogContent(req.Topics)
	if err != nil {
		log.Printf("Error generating blog content: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate content"})
		return
	}

	response := gin.H{"content": content}
	c.JSON(http.StatusOK, response)
}

// ReviewContentHandler handles requests to review content
func (bc *BlogController) ReviewContentHandler(c *gin.Context) {
	var req struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	suggestions, err := bc.blogUseCase.ReviewBlogContent(req.Content)
	if err != nil {
		log.Printf("Error generating suggestions: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to review content"})
		return
	}

	response := gin.H{"suggestions": suggestions}
	c.JSON(http.StatusOK, response)
}

// GenerateTopicHandler handles requests to generate topics
func (bc *BlogController) GenerateTopicHandler(c *gin.Context) {
	var req struct {
		Keywords []string `json:"keywords"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	topics, err := bc.blogUseCase.GenerateTrendingTopics(req.Keywords)
	if err != nil {
		log.Printf("Error generating trending topics: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate topics"})
		return
	}

	response := gin.H{"topics": topics}
	c.JSON(http.StatusOK, response)
}
