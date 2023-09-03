package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-jwt/initializers"
	"go-jwt/models"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	// Get token from header
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find user in database by ID saved in token subject
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// If user is found, set user in context
		c.Set("user", user)

		// Continue middleware processing
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// c.Next()
}
