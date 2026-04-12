package device

import "time"

type Device struct {
	ID         int64  `db:"id" json:"id"`
	DeviceCode string `db:"device_code" json:"device_code"`
	// UserId       int64  `db:"user_id" json:"user_id"`
	IsActive   bool       `db:"is_active" json:"is_active"`
	LastActive *time.Time `db:"last_active" json:"last_active"`
	ApiKeyHash string     `db:"apikey_hash" json:"-"`

	ActivationCode string `db:"activation_code" json:"-"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

//  --- DTOs

type DeviceRequest struct {
	DeviceCode     string `json:"device_code" validate:"required,max=16"`
	ActivationCode string `json:"activation_code" validate:"required,max=50"`
}

type DeviceClaimResponse struct {
	DeviceClaimId int64  `db:"id" json:"id"`
	DeviceId      int64  `db:"device_id" json:"device_id"`
	DeviceCode    string `db:"device_code" json:"device_code"`
	UserId        int64  `db:"user_id" json:"user_id"`
	DeviceName    string `db:"device_name" json:"device_name"`

	IsActive   bool       `db:"is_active" json:"is_active"`
	LastActive *time.Time `db:"last_active" json:"last_active"`

	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}
