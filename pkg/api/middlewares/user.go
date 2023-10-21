package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	// s := c.Request.Header.Get("Authorization")
	tokenString, err := c.Cookie("UserAuth")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	UsersId, err := ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("usersID", UsersId)
	c.Next()
}
