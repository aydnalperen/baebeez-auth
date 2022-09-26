package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(uid string, c *gin.Context) (string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["uid"] = uid
	claims["expiration_time"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	claims["ip"] = c.Request.Header.Get("X-Forwarded-For")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
func ExtractToken(ctx *gin.Context) string {

	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return bearerToken

}
func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)

	var keyFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// }
		// claims := token.Claims.(jwt.MapClaims)

		// if user_ip := claims["ip"]; user_ip != c.Request.Header.Get("X-Forwarded-For") {
		// 	return nil, fmt.Errorf("Different Ip : %v", claims["ip"])
		// }

		return []byte(os.Getenv("API_SECRET")), nil
	}
	_, err := jwt.Parse(tokenString, keyFunc)

	if err != nil {
		return err
	}
	return nil
}
func ExtractTokenUID(c *gin.Context) (string, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := fmt.Sprintf("%v", claims["uid"])
		return uid, nil
	}
	return "", nil
}
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
