package usecase

import (
	"context"
	"go-clean-arch/internal/adapter/repository"
	"go-clean-arch/internal/core/domain"
	"go-clean-arch/internal/utils"

	"go.uber.org/zap"
)

type UserService struct {
	UserRepo repository.UserRepository
	logger   *zap.SugaredLogger
}

func NewUserService(userRepo repository.UserRepository, logger *zap.SugaredLogger) *UserService {
	return &UserService{
		userRepo,
		logger,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		us.logger.Error("Failed to hash password: ", err)
		return domain.ErrInternal
	}
	user.Password = hashedPassword

	utils.BuildIdempotencyKey(user)

	err = us.UserRepo.Save(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			us.logger.Error("data already exist: ", err)
			return err
		}
		us.logger.Error("failed to create user: ", err)
		return domain.ErrInternal
	}

	return nil
}

func (us *UserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	user, err := us.UserRepo.Get(ctx, id)
	if err != nil {
		us.logger.Error("failed to get user: ", err)
		return nil, err
	}

	return user, nil
}

func (us *UserService) ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error) {
	users, err := us.UserRepo.List(ctx, skip, limit)
	if err != nil {
		us.logger.Error("failed to list users: ", err)
		return nil, err
	}

	return users, nil
}

func (us *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (us *UserService) DeleteUser(ctx context.Context, id string) error {
	return nil
}
