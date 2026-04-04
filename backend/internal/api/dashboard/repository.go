package dashboard

import (
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
		COALESCE(SUM(power_kwh), 0),
		COALESCE(AVG(voltage), 0),
		COALESCE(AVG(power_kwh), 0),
		COALESCE(AVG(current), 0)
	FROM energy_readings
	WHERE stamp >= date_trunc('month', $1)
	  AND stamp <  date_trunc('month', $1) + INTERVAL '1 month';
	`

	var result Overview

	err := r.db.QueryRow(query, month).Scan(
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

func (r *DashboardRepo) GetMonthlyAvgPower(year int) ([]MonthlyPower, error) {
	query := `
	SELECT
		m.month,
		COALESCE(AVG(e.power_kwh), 0) AS avg_power_draw
	FROM generate_series(
		DATE_TRUNC('year', $1::date),
		DATE_TRUNC('year', $1::date) + INTERVAL '11 months',
		INTERVAL '1 month'
	) AS m(month)
	LEFT JOIN energy_readings e
		ON date_trunc('month', e.stamp) = m.month
	GROUP BY m.month
	ORDER BY m.month;
	`

	// Convert year → time.Time (Jan 1 of that year)
	startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	rows, err := r.db.Query(query, startOfYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []MonthlyPower

	for rows.Next() {
		var m MonthlyPower
		if err := rows.Scan(&m.Month, &m.AvgPowerDraw); err != nil {
			return nil, err
		}
		results = append(results, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
