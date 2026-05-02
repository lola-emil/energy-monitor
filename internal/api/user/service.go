package user

import (
	"context"
	userModel "energy-monitor-server/internal/model/user"
	passwordutil "energy-monitor-server/internal/utils/password"
	"errors"
)

type UserService struct {
	repo userModel.UserRepository
}

func NewUserSerivce(repo userModel.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UpdateProfile(
	ctx context.Context,
	userID int64,
	req *UpdateProfileRequest,
) error {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if req.Username == "" {
		return errors.New("username is required")
	}

	// update password only if requested
	if req.NewPassword != "" {
		if req.CurrentPassword == "" {
			return errors.New("current password is required")
		}

		valid, err := passwordutil.Check(
			req.CurrentPassword,
			user.Password,
		)

		if err != nil {
			return err
		}

		if !valid {
			return errors.New("current password is incorrect")
		}

		hashedPassword, err := passwordutil.Hash(
			req.NewPassword,
			passwordutil.DefaultParams,
		)

		if err != nil {
			return err
		}

		user.Password = hashedPassword
	}

	user.Username = req.Username

	return s.repo.UpdateProfile(ctx, user)
}

func (s *UserService) GetProfile(
	ctx context.Context,
	userID int64,
) (*userModel.User, error) {
	return s.repo.GetByID(ctx, userID)
}
