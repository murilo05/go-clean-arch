package http

import (
	"go-clean-arch/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Document string `json:"document" binding:"required,min=11" example:"12345678911"`
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required" example:"murilo@gmail.com"`
	Age      int    `json:"age" binding:"required" example:"23"`
	Password string `json:"password" binding:"required,min=8" example:"12345678"`
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	create a new user account with default role "customer"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		registerRequest	true	"Register request"
//	@Success		200				{object}	userResponse	"User created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/users [post]
func (h *Handler) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	user := domain.User{
		Document: req.Document,
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		Password: req.Password,
	}

	err := h.userUseCase.CreateUser(ctx, &user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(&user)

	handleSuccess(ctx, rsp)
}
