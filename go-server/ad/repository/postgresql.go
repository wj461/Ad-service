package repository

import (
	"database/sql"

	"github.com/wj461/ad-service/swagger"
)

type postgresqlAdRepository struct {
	db *sql.DB
}

func NewPostgresqlAdRepository(db *sql.DB) *postgresqlAdRepository {
	return &postgresqlAdRepository{db: db}
}

func (p *postgresqlAdRepository) CreateAd(ad *swagger.Ad) error {

	return nil
}
