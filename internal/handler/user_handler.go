package handler

import (
	"errors"
	"strconv"

	"godest/internal/service"
	"godest/internal/transport/http/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(response.BadRequest(err.Error()))
		return
	}

	if err := h.svc.Register(req.Username, req.Email, req.Password); err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			_ = c.Error(response.Conflict(err.Error()))
			return
		}
		_ = c.Error(response.Internal("failed to register user", err))
		return
	}

	response.Created(c, gin.H{"message": "user registered successfully"})
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(response.BadRequest("invalid user ID"))
		return
	}

	user, err := h.svc.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			_ = c.Error(response.NotFound("user not found"))
			return
		}
		_ = c.Error(response.Internal("failed to get user", err))
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(response.BadRequest(err.Error()))
		return
	}

	resp, err := h.svc.Login(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) || errors.Is(err, service.ErrInvalidPassword) {
			_ = c.Error(response.Unauthorized("invalid username or password"))
			return
		}
		_ = c.Error(response.Internal("failed to login", err))
		return
	}

	response.Success(c, resp)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.svc.ListUsers()
	if err != nil {
		_ = c.Error(response.Internal("failed to get users", err))
		return
	}
	response.Success(c, users)
}
