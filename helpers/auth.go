package helpers

import (
	"api/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type authClaims struct {
	jwt.StandardClaims
	UserID string `json:"userId"`
}

func GenerateToken(user models.User) (string, error) {
	// expiresAt := time.Now().Add(24 * time.Hour).Unix()
	expiresAt := time.Now().Add(5 * 24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expiresAt,
		},
		UserID: user.ID,
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken(tokenString string) (string, string, error) {
	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		fmt.Printf("validateToken() err: %v\n\n", err)
		return "", "", err
	}

	if !token.Valid {
		return "", "", errors.New("invalid token")
	}

	id := claims.UserID
	email := claims.Subject
	return id, email, nil
}

func VerifyToken(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		fmt.Printf("VerifyToken() token: %v\n\n", token)
		fmt.Printf("VerifyToken() ok: %v\n\n", ok)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	id, email, err := validateToken(token)
	if err != nil {
		fmt.Printf("VerifyToken() err: %v\n\n", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	c.Set("id", id)
	c.Set("email", email)

	c.Writer.Header().Set("Authorization", "Bearer "+token)

	c.Next()
}

func getToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")

	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}

	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}

	return strings.Trim(arr[1], "\n\t\r"), true
}

func GetSession(c *gin.Context) (string, string, bool) {
	id, ok := c.Get("id")
	if !ok {
		return "", "", false
	}

	email, ok := c.Get("email")
	if !ok {
		return "", "", false
	}

	return id.(string), email.(string), true
}
