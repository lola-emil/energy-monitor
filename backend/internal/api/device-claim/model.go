package deviceclaim

import "time"

type DeviceClaim struct {
	ID         int64  `db:"id" json:"id"`
	DeviceId   int64  `db:"device_id" json:"device_id"`
	UserId     int64  `db:"user_id" json:"user_id"`
	DeviceName string `db:"device_name" json:"device_name"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}

type DeviceClaimRequest struct {
	DeviceCode string `json:"device_code" validate:"required,gt=0"`
	DeviceName string `json:"device_name"`
}
