package repository

import (
	"context"
	"go-clean-arch/internal/core/domain"

	"go.uber.org/zap"
)

type Repository struct {
	db     UserRepository
	logger *zap.SugaredLogger
}

func NewRepository(db UserRepository, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		db,
		logger,
	}
}

func (r *Repository) Save(ctx context.Context, user *domain.User) error {
	return r.db.Save(context.TODO(), user)
}

func (r *Repository) List(ctx context.Context, skip, limit uint64) ([]domain.User, error) {
	return r.db.List(context.TODO(), skip, limit)
}

func (r *Repository) Get(ctx context.Context, id string) (*domain.User, error) {
	return r.db.Get(context.TODO(), id)
}

func (r *Repository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	return r.db.Update(context.TODO(), user)
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	return r.db.Delete(context.TODO(), id)
}
