package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

func JWTMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractTokenFromHeader(c.Request)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		token, err := verifyJWTToken(tokenString, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("userID", token.Claims.(*CustomClaims).UserID)
		c.Next()
	}
}

func extractTokenFromHeader(req *http.Request) string {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}

	return strings.TrimSpace(splitToken[1])
}

func verifyJWTToken(tokenString, secretKey string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
}

type CustomClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}
