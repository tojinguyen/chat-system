package usecase

import (
	"context"
	"errors"

	"api-service/internal/domain"
	"api-service/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("User not found")
)

type UserUsecase interface {
	GetProfile(ctx context.Context, id uuid.UUID) (*domain.User, error)
	UpdateProfile(ctx context.Context, id uuid.UUID, name, phone, address *string) (*domain.User, error)
	SearchUsers(ctx context.Context, query string) ([]domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetProfile(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (u *userUsecase) UpdateProfile(ctx context.Context, id uuid.UUID, name, phone, address *string) (*domain.User, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	if name != nil && *name != "" {
		user.Name = *name
	}
	if phone != nil {
		user.Phone = phone
	}
	if address != nil {
		user.Address = address
	}

	if err := u.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) SearchUsers(ctx context.Context, query string) ([]domain.User, error) {
	if query == "" {
		return []domain.User{}, nil
	}
	return u.userRepo.Search(ctx, query)
}

func (u *userUsecase) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
