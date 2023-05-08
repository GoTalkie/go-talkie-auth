package handlers

import (
	"net/http"

	"github.com/GoTalkie/go-talkie-auth/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Config) Register(c *gin.Context) {
	var registerInput UserInfo

	if err := c.ShouldBindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Username: registerInput.Username,
		Password: registerInput.Password,
	}

	_, err := u.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

