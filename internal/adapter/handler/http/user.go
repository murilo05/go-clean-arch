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
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		registerRequest	true	"Register request"
//	@Success		200				{object}	userResponse	"User created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/user [post]
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

	handleCreated(ctx, rsp)
}

// listUsersRequest represents the request body for listing users
type listUsersRequest struct {
	Skip  uint64 `form:"skip" binding:"omitempty" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"5"`
}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	List users with pagination
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			skip	query		uint64			true	"Skip"
//	@Param			limit	query		uint64			true	"Limit"
//	@Success		200		{object}	meta			"User displayed"
//	@Failure		400		{object}	errorResponse	"Validation error"
//	@Failure		500		{object}	errorResponse	"Internal server error"
//	@Router			/user?skip=0&limit=5 [get]
//	@Security		BearerAuth
func (h *Handler) ListUsers(ctx *gin.Context) {
	var req listUsersRequest
	var usersList []userResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	users, err := h.userUseCase.ListUsers(ctx, req.Skip, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, user := range users {
		usersList = append(usersList, newUserResponse(&user))
	}

	total := uint64(len(usersList))
	meta := newMeta(total, req.Limit, req.Skip)
	rsp := toMap(meta, usersList, "users")

	handleSuccess(ctx, rsp)
}

func toMap(m meta, data any, key string) map[string]any {
	return map[string]any{
		"meta": m,
		key:    data,
	}
}

// getUserRequest represents the request body for getting a user
type getUserRequest struct {
	ID string `uri:"id" binding:"required,min=1" example:"1"`
}

// GetUser godoc
//
//	@Summary		Get a user
//	@Description	Get a user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"User ID"
//	@Success		200	{object}	userResponse	"User displayed"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/users/{id} [get]
//	@Security		BearerAuth
func (h *Handler) GetUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	user, err := h.userUseCase.GetUser(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(user)

	handleSuccess(ctx, rsp)
}
