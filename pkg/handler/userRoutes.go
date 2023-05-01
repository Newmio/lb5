package handler

import (
	"fmt"
	"lb5/pkg/models"
	userquery "lb5/pkg/postgres/userQuery"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := userquery.Post(h.DB, user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	c.IndentedJSON(http.StatusOK, models.Response{Response: fmt.Sprintf("user created, ligin: %s", user.Login)})
}

func (h *Handler) GetUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := userquery.GenerateToken(h.DB, user.Login, user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, models.Response{Response: token})
}
