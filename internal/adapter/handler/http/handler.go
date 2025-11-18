package http

import (
	"go-clean-arch/internal/core/usecase"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

type Handler struct {
	userUseCase usecase.UserUseCase
	//All useCases must be injected in the handler
}

func NewHTTPHandler(userUseCase usecase.UserUseCase) *Handler {
	return &Handler{
		userUseCase,
	}
}
