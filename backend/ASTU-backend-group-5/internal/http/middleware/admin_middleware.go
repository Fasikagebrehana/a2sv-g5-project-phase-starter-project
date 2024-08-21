package middleware

import (
	"blogApp/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found"})
			c.Abort()
			return
		}

		userClaims, ok := claims.(*domain.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			c.Abort()
			return
		}
		// fmt.Println(userClaims.Role, "role =============================")

		if userClaims.Role != "admin" && userClaims.Role != "owner" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		// If the user is an admin, proceed with the request
		c.Next()
	}
}
