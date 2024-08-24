package middleware

import (
	"blog_api/domain"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddlewareError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, domain.Response{"message": message})
	c.Abort()
}

/*
This is the authorization middleware used for the endpoints. It accepts a set of
roles for which the endpoint is open.

WORKFLOW:
  - Obtains the JWT from the authorization header
  - Parses the JWT and verifies the signature
  - Checks the role of the user associated with the token
  - Calls `c.Next()` if the querying user has permission to access the endpoint
*/
func AuthMiddlewareWithRoles(jwtService domain.JWTServiceInterface, cacheRepository domain.CacheRepositoryInterface, validRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// obtain token from the request header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			MiddlewareError(c, 401, "Authorization header not found")
			return
		}

		headerSegments := strings.Split(authHeader, " ")
		if len(headerSegments) != 2 || strings.ToLower(headerSegments[0]) != "bearer" {
			MiddlewareError(c, 401, "Authorization header is invalid")
			return
		}

		// check if the access token has been blacklisted
		// an access token is blacklisted when the user logs out
		if cacheRepository.IsCached(headerSegments[1]) {
			MiddlewareError(c, 401, "User has already logged out")
			return
		}

		// parses token with the correct signing method and checks for errors and token validity
		token, validErr := jwtService.ValidateAndParseToken(headerSegments[1])
		if validErr != nil {
			MiddlewareError(c, 401, validErr.Error())
			return
		}

		// check whether the token is an accessToken
		tokenType, err := jwtService.GetTokenType(token)
		if err != nil {
			MiddlewareError(c, 401, err.Error())
			return
		}

		if tokenType != "accessToken" {
			MiddlewareError(c, 401, "Invalid token type: make sure to use the accessToken to authorize actions")
			return
		}

		// check the expiry date of the token
		expiresAtTime, err := jwtService.GetExpiryDate(token)
		if err != nil {
			MiddlewareError(c, 401, err.Error())
			return
		}

		if expiresAtTime.Compare(time.Now()) == -1 {
			MiddlewareError(c, 401, "Token expired")
			return
		}

		// get the role from the claims of the JWT
		userRole, err := jwtService.GetRole(token)
		if err != nil {
			MiddlewareError(c, 401, err.Error())
			return
		}

		// get the username from the claims of the JWT
		username, err := jwtService.GetUsername(token)
		if err != nil {
			MiddlewareError(c, 401, err.Error())
			return
		}

		valid := false
		for _, validRole := range validRoles {
			if userRole == validRole {
				valid = true
				break
			}
		}

		if !valid {
			MiddlewareError(c, 403, fmt.Sprintf("'%v' roles are not allowed to access this endpoint", userRole))
			return
		}

		// pass the username to the handler
		c.Set("username", username)
		c.Next()
	}
}
