package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GoTalkie/go-talkie-auth/internal/models"
	"github.com/GoTalkie/go-talkie-auth/internal/utils/token"
)

func (conf *Config) User(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
