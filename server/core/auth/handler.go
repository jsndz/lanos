package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

type RegisterRequest struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Register(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.UserService.Register(c.Request.Context(),
		req.UserName,
		req.Email,
		req.Password,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.SetCookie("session_token", user.SessionToken, 7*24*3600, "/", "", true, true)
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.User.ID,
		"username": user.User.Username,
		"email":    user.User.Email,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ip := c.ClientIP()
	user, err := h.UserService.Login(c.Request.Context(), req.Email, req.Password, ip)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.SetCookie("session_token", user.SessionToken, 7*24*3600, "/", "", true, true)

	c.JSON(http.StatusCreated, gin.H{
		"username":      user.User.Username,
		"email":         user.User.Email,
		"session_token": user.SessionToken,
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	token, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Session token not found",
		})
		return
	}

	success, err := h.UserService.Logout(c.Request.Context(), token)
	if err != nil || !success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to logout",
		})
		return
	}

	c.SetCookie("session_token", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
