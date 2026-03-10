package mymqtt

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetDeviceByIdAndSerial(id int64, serial string) (*Device, error) {
	query := "SELECT * FROM devices WHERE id = $1 AND device_serial = $2"

	var device Device
	if err := r.db.Get(&device, query, id, serial); err != nil {
		return nil, err
	}

	return &device, nil
}
