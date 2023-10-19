package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/msssrp/SportEquipmentBorrowing/function"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type CustomClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret, err := function.GetDotEnv("SECRET")
		if err != nil {
			panic(err)
		}

		tokenString := extractTokenFromHeader(c.Request)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		token, err := verifyJWTToken(tokenString, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			c.Set("userID", claims.UserID)
		} else {
			spew.Println("Invalid token")
		}

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

func verifyJWTToken(tokenString string, secretKey string) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
