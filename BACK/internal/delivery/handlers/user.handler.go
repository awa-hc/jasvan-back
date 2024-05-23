package handlers

import (
	"jasvan/internal/domain/entities"
	"jasvan/internal/domain/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var User entities.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.CreateUser(c, &User); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": User})
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.UserService.GetUserByEmail(c, email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := h.UserService.GetUserByUsername(c, username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	users, err := h.UserService.GetAllUser(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": users})
}
