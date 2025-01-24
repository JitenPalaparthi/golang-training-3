package handlers

import (
	"demo/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("demotoken") // vault
// user and password
// token
// token

func Ping(ctx *gin.Context) {
	ctx.String(200, "Pong")
	ctx.Abort()
}

func Health(ctx *gin.Context) {
	//ctx.JSON(http.StatusOK, map[string]string{"health": "ok"})
	ctx.JSON(http.StatusOK, gin.H{"health": "ok"})
	ctx.Abort()
}

func Login(ctx *gin.Context) {
	userLogin := new(models.UserLogin)
	err := ctx.Bind(userLogin)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "invalid user data")
		ctx.Abort()
	}
	err = userLogin.Validate()
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	// write logic to validate username and passowrd from the database

	if userLogin.Email == "jitenp@outlook.com" && userLogin.Passwortd == "jiten123" {
		println("-------->>>>>>>>>>>")
		token, err := GenerateJWT(userLogin.Email)
		if err != nil {
			log.Println(err.Error())
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
		}

		ctx.String(200, token)
		// How do you generate the token ?
	}
}

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Token is valid; extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("email", claims["email"])
		}
		c.Next()
	}
}

func Restrict(c *gin.Context) {
	log.Println(c.Request.URL)
	log.Println(c.RemoteIP())
	isValid := false

	if isValid {
		c.Next()
	} else {
		c.String(401, "restricted user")
		c.Abort()
		return
	}
}
