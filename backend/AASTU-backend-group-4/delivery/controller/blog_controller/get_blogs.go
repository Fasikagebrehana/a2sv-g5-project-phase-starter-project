package blog_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (bc *BlogController) GetBlogs(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	sortBy := c.DefaultQuery("sort_by", "recent") // default to "recent"

	posts, totalPosts, err := bc.usecase.GetBlogs(c, page, limit, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (totalPosts + limit - 1) / limit

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"pagination": gin.H{
			"current_page": page,
			"total_pages":  totalPages,
			"total_posts":  totalPosts,
		},
	})
}
