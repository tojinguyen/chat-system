package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"api-service/internal/config"
	"api-service/internal/domain"
	"api-service/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists   = errors.New("Username already exists")
	ErrInvalidCredentials  = errors.New("Invalid credentials")
	ErrUnauthorized        = errors.New("Unauthorized")
)

type AuthTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthClaims struct {
	UserID   string `json:"sub"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthUsecase interface {
	Register(ctx context.Context, username, password, name string, phone, address *string) (*domain.User, error)
	Login(ctx context.Context, username, password string) (*AuthTokens, error)
	ValidateToken(tokenStr string) (*AuthClaims, error)
}

type authUsecase struct {
	cfg      *config.Config
	userRepo repository.UserRepository
}

func NewAuthUsecase(cfg *config.Config, userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (u *authUsecase) Register(ctx context.Context, username, password, name string, phone, address *string) (*domain.User, error) {
	existing, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := &domain.User{
		ID:           uuid.New(),
		Username:     username,
		PasswordHash: string(hash),
		Name:         name,
		Phone:        phone,
		Address:      address,
	}

	if err := u.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *authUsecase) Login(ctx context.Context, username, password string) (*AuthTokens, error) {
	user, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return u.generateTokens(user.ID, user.Username)
}

func (u *authUsecase) generateTokens(userID uuid.UUID, username string) (*AuthTokens, error) {
	now := time.Now()
	accessExpiry := now.Add(u.cfg.JWTExpiresIn)
	refreshExpiry := now.Add(7 * 24 * time.Hour)

	accessClaims := AuthClaims{
		UserID:   userID.String(),
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	refreshClaims := AuthClaims{
		UserID:   userID.String(),
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiry),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &AuthTokens{
		AccessToken:  accessStr,
		RefreshToken: refreshStr,
	}, nil
}

func (u *authUsecase) ValidateToken(tokenStr string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(u.cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, ErrUnauthorized
	}

	claims, ok := token.Claims.(*AuthClaims)
	if !ok || !token.Valid {
		return nil, ErrUnauthorized
	}

	return claims, nil
}
