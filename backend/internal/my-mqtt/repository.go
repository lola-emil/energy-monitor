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

func (r *Repository) SaveDeviceReadings(data EnergyReadingBody) (int64, error) {
	query := `
	INSERT INTO energy_readings 
	(device_id, voltage, current, power_kwh)
	VALUES
	($1, $2, $3, $4)
	RETURNING id
	`

	var id int64
	err := r.db.QueryRow(query,
		data.DeviceId,
		data.Voltage,
		data.Current, data.PowerKwh,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
