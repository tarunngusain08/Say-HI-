package auth

import (
	"Say-Hi/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

var blacklist = make(map[string]bool)

// When a user logs out, add their token to the blacklist
func addToBlacklist(token string) {
	blacklist[token] = true
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(config.Authorization)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{config.Error: config.Unauthorized})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.SecretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{config.Error: config.Unauthorized})
			c.Abort()
			return
		}

		// Set user information in the context for further use
		claims := token.Claims.(jwt.MapClaims)
		c.Set(config.UserId, claims[config.UserId])
		c.Next()
	}
}
