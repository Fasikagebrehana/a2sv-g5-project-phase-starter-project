package middlewares

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function that handles authentication for incoming requests.
// It checks the Authorization header for a valid access token and refresh token.
// If the access token is valid, it extracts the claims and sets them in the context.
// If the access token is invalid or expired, it checks the refresh token.
// If the refresh token is valid, it issues new access and refresh tokens.
// If the refresh token is invalid or expired, it returns an error response.
// The middleware function requires an instance of bootstrap.Env and domain.TokenInfrastructure as parameters.
// It returns a gin.HandlerFunc that can be used as middleware in a Gin router.
func AuthMiddleware(env *bootstrap.Env, tokenService domain.TokenInfrastructure) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		tokens := strings.Split(authParts[1], ":")
		if len(tokens) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Both access and refresh tokens are required"})
			c.Abort()
			return
		}

		accessToken := tokens[0]
		refreshToken := tokens[1]

		// Validate access token
		ok, _ := tokenService.ValidateToken(accessToken, env.AccessTokenSecret)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}

		ok, _ = tokenService.CheckTokenExpiry(accessToken, env.AccessTokenSecret)

		if !ok {

			// Access token is invalid or expired, check the refresh token
			_, refreshErr := tokenService.CheckTokenExpiry(refreshToken, env.RefreshTokenSecret)
			if refreshErr != nil {
				log.Printf("Refresh token expired or invalid: %v", refreshErr)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired or invalid, please login again"})
				c.Abort()
				return
			}

			// Refresh token is valid, issue new access and refresh tokens
			newClaims, claimsErr := tokenService.ExtractClaims(accessToken, env.AccessTokenSecret)
			if claimsErr != nil {
				log.Printf("Error extracting claims: %v", claimsErr)
				c.JSON(http.StatusUnauthorized, gin.H{"error": claimsErr.Error()})
				c.Abort()
				return
			}

			accessToken, refreshToken, updateErr := tokenService.UpdateTokens(newClaims["UserID"].(string))
			if updateErr != nil {
				log.Printf("Error updating tokens: %v", updateErr)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to update tokens"})
				c.Abort()
				return
			}

			log.Printf("Tokens updated successfully: accessToken: %v, refreshToken: %v", accessToken, refreshToken)
			c.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
			c.Abort()
			return
		}

		claims, err := tokenService.ExtractClaims(accessToken, env.AccessTokenSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Access token is valid
		c.Set("userID", claims["UserID"])
		c.Set("userRole", claims["Role"])
		c.Set("claims", claims)
		c.Next()
	}
}
