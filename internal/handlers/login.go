package handlers

import (
	"net/http"

	"github.com/GoTalkie/go-talkie-auth/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Config) Login(c *gin.Context) {
	var loginInput UserInfo

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Username: loginInput.Username,
		Password: loginInput.Password,
	}

	token, err := models.LoginCheck(u.Username, u.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
