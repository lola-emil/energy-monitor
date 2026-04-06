package deviceclaim

import "time"

type DeviceClaim struct {
	ID       int64 `db:"id" json:"id"`
	DeviceId int64 `db:"device_id" json:"device_id"`
	UserId   int64 `db:"user_id" json:"user_id"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}

type DeviceClaimRequest struct {
	DeviceId int64 `json:"device_id" validate:"required,gt=0"`
	UserId   int64 `json:"user_id" validate:"required,gt=0"`
}
