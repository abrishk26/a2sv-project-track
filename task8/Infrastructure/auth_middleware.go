package infrastructures

import (
	"net/http"
	"strings"

	usecases "github.com/abrishk26/a2sv-project-track/task8/Usecases"
	"github.com/gin-gonic/gin"
)

func IsLoggedIn(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header missing",
		})
		c.Abort()
		return
	}

	split := strings.Split(authHeader, " ")
	if len(split) < 2 || strings.ToLower(split[0]) != "bearer" || split[1] == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token format",
		})
		c.Abort()
		return
	}

	tokenString := split[1]

	ctx := usecases.ContextWithToken(c.Request.Context(), tokenString)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
