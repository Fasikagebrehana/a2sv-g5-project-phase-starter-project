package controllers

import (
	domain "blogs/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleResponse handles the API response based on the type of the response object.
func HandleResponse(c *gin.Context, response interface{}) {

	switch res := response.(type) {
	case *domain.SuccessResponse:
		c.JSON(http.StatusOK, res)
	case *domain.ErrorResponse:
		c.JSON(res.Status, res)
	case *domain.URL:
		url := res.URL
		c.Redirect(http.StatusTemporaryRedirect, url)
	case *domain.LoginResponse:
		c.JSON(http.StatusOK, res)
	case *domain.AiResponse:
		c.JSON(http.StatusOK , res)
	case *domain.UserPromotionResponse:
		c.JSON(http.StatusOK , res)
	// case *domain.TaskResponse:
	// 	c.JSON(http.StatusOK , res)
	// case *domain.TaskSuccessResponse:
	// 	c.JSON(http.StatusOK , res)
	// case *domain.SingleTaskResponse:
	// 	c.JSON(http.StatusOK , res)
	default:
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Internal Server Error", Status: 500})
	}
}
