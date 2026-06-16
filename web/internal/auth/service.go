package auth

import (
	"context"
	"energy-monitor-server/internal/model/user"
	jwtutil "energy-monitor-server/internal/utils/jwt"
	passwordutil "energy-monitor-server/internal/utils/password"
	"errors"
	"log"
)

type AuthService struct {
	userRepo user.UserRepository
}

func NewAuthService(userRepo user.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		log.Println(err.Error())
		return "", ErrInvalidCredentials
	}

	if !user.IsActive {
		return "", ErrAccountDisabled
	}

	passwordMatched, err := passwordutil.Check(password, user.Password)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	if !passwordMatched {
		return "", ErrInvalidCredentials
	}

	token, err := jwtutil.Generate(user.ID, string(user.Role))
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	_ = s.userRepo.UpdateLastLogin(ctx, user.ID)

	return token, nil
}

func (s *AuthService) Register(ctx context.Context, username, name, plainPassword string) (string, error) {
	//  Check if user already exists
	existing, _ := s.userRepo.GetByUsername(ctx, username)
	if existing != nil {
		return "", errors.New("user taken")
	}

	// Hash password
	hashed, err := passwordutil.Hash(plainPassword, passwordutil.DefaultParams)
	if err != nil {
		return "", err
	}

	// Create user
	user := &user.User{
		Username: username,
		Password: hashed,
		Name:     name,
		IsActive: true,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return "", err
	}

	// Generate JWT (auto-login after register)
	token, err := jwtutil.Generate(user.ID, string(user.Role))
	if err != nil {
		return "", err
	}

	return token, nil
}
