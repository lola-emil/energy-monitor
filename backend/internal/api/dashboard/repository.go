package dashboard

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type DashboardRepo struct {
	db *sqlx.DB
}

func NewDashboardRepo(db *sqlx.DB) *DashboardRepo {
	return &DashboardRepo{
		db: db,
	}
}

func (r *DashboardRepo) GetOverview(month time.Time) (*Overview, error) {
	query := `
	SELECT
		COALESCE(SUM(power * 2.0 / 3600.0 / 1000.0), 0),
		COALESCE(AVG(voltage), 0),
		COALESCE(AVG(power) / 1000, 0),
		COALESCE(AVG(current), 0)
	FROM readings_raw
	WHERE bucket >= $1 AND bucket < $2;
	`

	start := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	var result Overview

	err := r.db.QueryRow(query, start, end).Scan(
		&result.TotalEnergyConsumed,
		&result.AvgVoltage,
		&result.AvgPowerDraw,
		&result.AvgCurrent,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *DashboardRepo) GetMonthyEnergyConsumption(ctx context.Context) ([]MonthlyConsumption, error) {
	query := `
	SELECT
		date_trunc('month', bucket) AS month,
		SUM(power * 2.0 / 3600.0 / 1000.0) AS energy_kwh
	FROM readings_raw
	GROUP BY month
	ORDER BY month;
	`
	var result []MonthlyConsumption

	if err := r.db.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}

	return result, nil
}
