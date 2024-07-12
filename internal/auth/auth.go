package auth

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hculpan/osetools/internal/dbutils"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("missing .env file: %s", err.Error())
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not defined")
	}

	jwtKey = []byte(secret)
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token for a given username
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func GetCurrentUsername(r *http.Request) string {
	c, err := r.Cookie("token")
	if err != nil {
		return ""
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return ""
	}

	if claims, ok := tkn.Claims.(*Claims); ok && tkn.Valid {
		return claims.Username
	} else {
		return ""
	}
}

func IsAuthorized(r *http.Request) bool {
	c, err := r.Cookie("token")
	if err != nil {
		return false
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false
	}

	if !tkn.Valid {
		return false
	}

	return true
}

func ValidateUser(username, password string) bool {
	if username == "" || password == "" {
		return false
	}

	hashedPassword, err := dbutils.RetrieveUserHash(username)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == nil {
		slog.Info("user logged in", "username", username)
	}

	return err == nil
}
