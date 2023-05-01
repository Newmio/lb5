package handler

import (
	"errors"
	userquery "lb5/pkg/postgres/userQuery"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	userctx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.IndentedJSON(http.StatusUnauthorized, "header = null")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.IndentedJSON(http.StatusUnauthorized, "invalid header")
		return
	}

	userId, err := userquery.ParseToken(headerParts[1])
	if err != nil {
		panic(err)
	}

	c.Set(userctx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userctx)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}
