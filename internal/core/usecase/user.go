package usecase

import (
	"context"
	"go-clean-arch/internal/core/domain"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, id uint64) (*domain.User, error)
	ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint64) error
}
