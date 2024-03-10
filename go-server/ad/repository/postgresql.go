package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/wj461/ad-service/domain"
	"github.com/wj461/ad-service/swagger"
)

type postgresqlAdRepository struct {
	db *sqlx.DB
}

func NewPostgresqlAdRepository(db *sqlx.DB) domain.AdRepository {
	return &postgresqlAdRepository{db: db}
}

func (p *postgresqlAdRepository) CreateAd(ctx context.Context, ad *swagger.Ad) error {
	sqlStatement := `
	INSERT INTO ad (title, start_at, end_at, age_start, age_end, male, female) VALUES 
	($1, $2, $3, $4, $5,$6, $7)
	`

	tx := p.db.MustBegin()

	tx.MustExec(sqlStatement, ad.Title, ad.StartAt, ad.EndAt, ad.Conditions.AgeStart, ad.Conditions.AgeEnd, ad.Conditions.Male, ad.Conditions.Female)

	tx.Commit()

	var ads []swagger.Ad

	p.db.Select(&ads, sqlStatement)

	return nil
}
