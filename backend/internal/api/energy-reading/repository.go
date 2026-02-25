package energyreading

import "github.com/jmoiron/sqlx"

type EnergyReadingRepo struct {
	db *sqlx.DB
}

func NewEnergyReadingRepo(db *sqlx.DB) *EnergyReadingRepo {
	return &EnergyReadingRepo{
		db: db,
	}
}
