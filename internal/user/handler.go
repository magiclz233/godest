package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler 用户 HTTP 入口
type Handler struct {
	svc *Service
}

// NewHandler 创建用户 Handler
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Register 用户注册
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Register(req.Username, req.Email, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "用户注册成功"})
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 用户登录
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.svc.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListUsers 获取用户列表
func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.svc.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户列表"})
		return
	}
	c.JSON(http.StatusOK, users)
}
