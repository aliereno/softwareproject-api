package auth

import (
	"github.com/aliereno/softwareproject-api/internal/orm/models"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

// Ping is simple keep-alive/ping handler
func GetToken(user models.User) (string, error) {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"role": user.Role,
		"exp":  time.Now().Add(time.Minute * 1800).Unix(), // TOKEN EXPIRE TIME
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mysupersecretpassword))
	if err != nil {
		return "Could not generate token", err
	}
	return tokenString, nil
}

func LookUserTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := []byte(mysupersecretpassword)
			return b, nil
		})
		if err != nil {
			_ = c.AbortWithError(401, err)
		}
	}
}

func LookAdminTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := []byte(mysupersecretpassword)
			return b, nil
		})
		if err != nil {
			_ = c.AbortWithError(401, err)
			return
		}
		if claims, ok := token.Claims.(jwt_lib.MapClaims); ok && token.Valid {
			if claims["role"] != 1 {
				_ = c.AbortWithError(401, err)
				return
			}
		}
	}
}
func LookSupportTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := []byte(mysupersecretpassword)
			return b, nil
		})
		if err != nil {
			_ = c.AbortWithError(401, err)
			return
		}
		if claims, ok := token.Claims.(jwt_lib.MapClaims); ok && token.Valid {
			if claims["role"] == 0 {
				_ = c.AbortWithError(401, err)
				return
			}
		}
	}
}
