export interface Settings {
  id?: number
  user_id?: number

  currency: string
  rate_per_kwh: number
  fixed_monthly_charge: number

  default_analytics_range: string
  refresh_interval_seconds: number
  time_format: string

  enable_voltage_alerts: boolean
  over_voltage_threshold: number
  under_voltage_threshold: number

  enable_current_alerts: boolean
  over_current_threshold: number

  enable_offline_alerts: boolean

  updated_at?: string
}