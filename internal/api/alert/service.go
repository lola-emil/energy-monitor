package alertapi

import (
	"context"
	"energy-monitor-server/internal/model/alert"
)

type AlertService struct {
	repo alert.AlertRepository
}

func NewAlertService(repo alert.AlertRepository) *AlertService {
	return &AlertService{repo: repo}
}

func (s *AlertService) List(ctx context.Context, userID int64, f alert.AlertFilter) ([]alert.Alert, error) {
	return s.repo.List(ctx, userID, f)
}

func (s *AlertService) Resolve(ctx context.Context, userID, id int64) error {
	return s.repo.Resolve(ctx, userID, id)
}

func (s *AlertService) Get(ctx context.Context, userID, id int64) (*alert.Alert, error) {
	return s.repo.GetByID(ctx, userID, id)
}

func (s *AlertService) GetAnalyticsAlerts(
	ctx context.Context,
	userID int64,
	applianceID *int64,
	rangeType string,
) ([]alert.Alert, error) {

	return s.repo.GetAnalyticsAlerts(
		ctx,
		userID,
		applianceID,
		rangeType,
		5,
	)
}

func (s *AlertService) GetRecentByAppliance(
	ctx context.Context,
	userID int64,
	applianceID int64,
) ([]alert.Alert, error) {

	return s.repo.GetRecentByAppliance(
		ctx,
		userID,
		applianceID,
		5,
	)
}
