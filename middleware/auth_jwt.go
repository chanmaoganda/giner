package middleware

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	publicKeyPath = "./certs/public.pem"
	privateKeyPath = "./certs/private.pem"
	keyLoadOnce sync.Once
	publicKey *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func loadKeys() {
	keyLoadOnce.Do(
		func() {
			log.Println("This message should only be once")

			publicPem, err := os.ReadFile(publicKeyPath)
			if err != nil {
				log.Fatalln("No Public Keys")
			}

			publicKey, _ = jwt.ParseRSAPublicKeyFromPEM(publicPem)

			privatePem, err := os.ReadFile(privateKeyPath)
			if err != nil {
				log.Fatalln("No Private Keys")
			}

			privateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(privatePem)
		})
}

func SignToken(username string) (string, error) {
	loadKeys()

	signTime := time.Now()
	expireTime := signTime.Add(time.Hour * 12)
	
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(signTime),
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)

    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func AuthMiddleWareJWT() gin.HandlerFunc {
	loadKeys()
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Missing Auth Header",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Header type incorrect, requiring Bearer header",
			})
			c.Abort()
			return
		}

		claims := &Claims{}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
		
				return publicKey, nil
			})

		if err != nil {
			log.Println("[ERROR] parse claims error ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
