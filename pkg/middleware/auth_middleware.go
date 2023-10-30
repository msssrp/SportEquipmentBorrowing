package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

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

func AccessTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret, err := function.GetDotEnv("SECRET")
		if err != nil {
			panic(err)
		}
		// Extract the token from the Authorization header
		accessToken := extractTokenFromHeader(c.Request)

		// Check if the access token is expired
		if isAccessTokenExpired(accessToken, secret) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort() // Abort the request, preventing further processing
			return
		}

		// Call the next middleware or handler
		c.Next()
	}
}

func isAccessTokenExpired(tokenString string, secretKey string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method and provide the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	// Check for parsing errors
	if err != nil {
		return true
	}

	// Check if the token is valid and not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		return expirationTime.Before(time.Now())
	}

	return true
}

func AuthenticateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret, err := function.GetDotEnv("SESSION_SECRET")
		if err != nil {
			panic(err)
		}
		shortSessionToken := c.GetHeader("sessionToken")
		if shortSessionToken == "" {
			shortSessionToken = c.Query("sessionToken")
		}

		// Verify the short session token
		token, err := jwt.Parse(shortSessionToken, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// You need to provide your own secret key for decoding the short session token
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token"})
			c.Abort()
			return
		}

		// Token is valid, proceed to the next middleware or handler
		c.Next()
	}
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret, err := function.GetDotEnv("SECRET")
		if err != nil {
			panic(err)
		}
		session_secret, err := function.GetDotEnv("SESSION_SECRET")
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

		shortSessionToken, err := jwt.New(jwt.SigningMethodHS256).SignedString([]byte(session_secret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short session token"})
			return
		}

		c.Set("sessionToken", shortSessionToken)

		c.Next()
	}
}

func JWTGetClaims() gin.HandlerFunc {
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

func JWTVerify() gin.HandlerFunc {
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

		c.Set("token", token)

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
