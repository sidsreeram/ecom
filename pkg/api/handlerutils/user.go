package handlerutils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(c *gin.Context) (int, error) {
	id, err := c.Cookie("userId")
	if err != nil {
		return 0, err
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
