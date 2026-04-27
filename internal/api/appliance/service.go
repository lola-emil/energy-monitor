package applianceapi

import (
	"context"
	"energy-monitor-server/internal/model/appliance"
	"errors"
	"strings"
)

type ApplianceService struct {
	repo appliance.ApplianceRepository
}

func NewApplianceService(r appliance.ApplianceRepository) *ApplianceService {
	return &ApplianceService{repo: r}
}

func (s *ApplianceService) Create(
	ctx context.Context,
	a *appliance.Appliance,
) error {
	if a.Name == "" {
		return errors.New("name is required")
	}

	if a.Location == "" {
		return errors.New("location is required")
	}

	if a.DeviceCode == "" {
		return errors.New("device code is required")
	}

	if !strings.HasPrefix(a.DeviceCode, "EMS-") {
		return errors.New("invalid device code")
	}

	a.Status = appliance.ApplianceStatusOffline

	return s.repo.Create(ctx, a)
}

func (s *ApplianceService) List(ctx context.Context, userID int64) ([]appliance.Appliance, error) {
	return s.repo.List(ctx, userID)
}

func (s *ApplianceService) Get(ctx context.Context, userID, id int64) (*appliance.Appliance, error) {
	return s.repo.GetByID(ctx, userID, id)
}

func (s *ApplianceService) Update(ctx context.Context, a *appliance.Appliance) error {
	return s.repo.Update(ctx, a)
}

func (s *ApplianceService) Delete(ctx context.Context, userID, id int64) error {
	return s.repo.Delete(ctx, userID, id)
}
