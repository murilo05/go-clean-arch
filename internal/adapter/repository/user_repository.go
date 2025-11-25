package repository

import (
	"context"
	"go-clean-arch/internal/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	List(ctx context.Context, skip, limit uint64) ([]domain.User, error)
	Get(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}
