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
//	@Success		201				{object}	userResponse	"User created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/v1/user [post]
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
//	@Param			skip	query		uint64			false	"Skip"
//	@Param			limit	query		uint64			true	"Limit"
//	@Success		200		{object}	response		"Users listed successfully"
//	@Failure		400		{object}	errorResponse	"Validation error"
//	@Failure		500		{object}	errorResponse	"Internal server error"
//	@Router			/v1/user [get]
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
	ID string `uri:"id" binding:"required,min=1" example:"f99c44eb088fbc06a040a359491b19ac479deca49b84508c9524eb41463a14dd"`
}

// GetUser godoc
//
//	@Summary		Get a user
//	@Description	Get a user by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"User ID"
//	@Success		200	{object}	response		"User displayed successfully"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/v1/user/{id} [get]
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

// updateUserRequest represents the request body to update a user
type updateUserRequest struct {
	ID       string `json:"id" binding:"required" example:"f99c44eb088fbc06a040a359491b19ac479deca49b84508c9524eb41463a14dd"`
	Document string `json:"document" binding:"required" example:"12345678911"`
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required" example:"murilo@gmail.com"`
	Age      int    `json:"age" binding:"required" example:"23"`
	Password string `json:"password" binding:"required" example:"12345678"`
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Update a user by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			updateUserRequest	body		updateUserRequest	true	"Update user request"
//	@Success		200	{object}	response		"User updated successfully"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/v1/user [put]
//	@Security		BearerAuth
func (h *Handler) UpdateUser(ctx *gin.Context) {

	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	user := domain.User{
		ID:       req.ID,
		Document: req.Document,
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		Password: req.Password,
	}

	err := h.userUseCase.UpdateUser(ctx, &user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(&user)

	handleUpdated(ctx, rsp)
}

// deleteUserRequest represents the request body to delete a user
type deleteUserRequest struct {
	ID string `uri:"id" binding:"required,min=1" example:"f99c44eb088fbc06a040a359491b19ac479deca49b84508c9524eb41463a14dd"`
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"User ID"
//	@Success		200	{object}	response		"User deleted successfully"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/v1/user/{id} [delete]
//	@Security		BearerAuth
func (h *Handler) DeleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	err := h.userUseCase.DeleteUser(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleDeleted(ctx, req.ID)
}
