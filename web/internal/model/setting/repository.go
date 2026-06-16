package setting

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type settingsRepo struct {
	db *sqlx.DB
}

type SettingsRepository interface {
	GetByUserID(ctx context.Context, userID int64) (*Settings, error)
	Upsert(
		ctx context.Context,
		s *Settings,
	) error
}

func NewSettingsRepository(db *sqlx.DB) SettingsRepository {
	return &settingsRepo{db: db}
}

func (r *settingsRepo) GetByUserID(ctx context.Context, userID int64) (*Settings, error) {
	var s Settings

	err := r.db.GetContext(ctx, &s, `
		SELECT *
		FROM settings
		WHERE user_id = $1
		LIMIT 1
	`, userID)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *settingsRepo) Upsert(
	ctx context.Context,
	s *Settings,
) error {
	query := `
		INSERT INTO settings (
			user_id,

			currency,
			rate_per_kwh,
			fixed_monthly_charge,

			default_analytics_range,
			refresh_interval_seconds,
			time_format,

			enable_voltage_alerts,
			over_voltage_threshold,
			under_voltage_threshold,

			enable_current_alerts,
			over_current_threshold,

			enable_offline_alerts,

			updated_at
		)
		VALUES (
			$1, $2, $3, $4,
			$5, $6, $7,
			$8, $9, $10,
			$11, $12,
			$13,
			NOW()
		)
		ON CONFLICT (user_id)
		DO UPDATE SET

			currency = EXCLUDED.currency,
			rate_per_kwh = EXCLUDED.rate_per_kwh,
			fixed_monthly_charge = EXCLUDED.fixed_monthly_charge,

			default_analytics_range = EXCLUDED.default_analytics_range,
			refresh_interval_seconds = EXCLUDED.refresh_interval_seconds,
			time_format = EXCLUDED.time_format,

			enable_voltage_alerts = EXCLUDED.enable_voltage_alerts,
			over_voltage_threshold = EXCLUDED.over_voltage_threshold,
			under_voltage_threshold = EXCLUDED.under_voltage_threshold,

			enable_current_alerts = EXCLUDED.enable_current_alerts,
			over_current_threshold = EXCLUDED.over_current_threshold,

			enable_offline_alerts = EXCLUDED.enable_offline_alerts,

			updated_at = NOW()
	`

	_, err := r.db.ExecContext(ctx, query,
		s.UserID,

		s.Currency,
		s.RatePerKWh,
		s.FixedMonthlyCharge,

		s.DefaultAnalyticsRange,
		s.RefreshIntervalSeconds,
		s.TimeFormat,

		s.EnableVoltageAlerts,
		s.OverVoltageThreshold,
		s.UnderVoltageThreshold,

		s.EnableCurrentAlerts,
		s.OverCurrentThreshold,

		s.EnableOfflineAlerts,
	)

	return err
}
