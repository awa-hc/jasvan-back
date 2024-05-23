package handlers

import (
	"jasvan/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) LoginWithEmail(c *gin.Context) {
	var login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthService.LoginWithEmail(c, login.Email, login.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(200, gin.H{"token": token})

}

func (h *AuthHandler) GetUserByContext(c *gin.Context) {
	user, err := h.AuthService.GetUserByContext(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})
}
