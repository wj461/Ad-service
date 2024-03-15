package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

	tx := p.db.MustBegin()
	//insert a row into ad table
	sqlStatement := `
	INSERT INTO ad (title, start_at, end_at, age_start, age_end, gender) VALUES 
	($1, $2, $3, $4, $5, $6) RETURNING ad_id
	`
	ad_id := 0
	err := tx.QueryRowx(sqlStatement, ad.Title, ad.StartAt, ad.EndAt, ad.Conditions.AgeStart, ad.Conditions.AgeEnd, ad.Conditions.Gender).Scan(&ad_id)
	if err != nil {
		logrus.Error(err)
		return err
	}
	sqlStatement = `
	INSERT INTO country (ad_id, country_name) VALUES 
	($1, $2)
	`
	for _, country := range ad.Conditions.Country {
		_, err = tx.Exec(sqlStatement, ad_id, country)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	sqlStatement = `
	INSERT INTO platform (ad_id, platform_name) VALUES 
	($1, $2)
	`
	for _, platform := range ad.Conditions.Platform {
		_, err = tx.Exec(sqlStatement, ad_id, platform)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	defer tx.Rollback()

	tx.Commit()

	return nil
}

func (p *postgresqlAdRepository) SearchAd(ctx context.Context, adQuery *swagger.AdQuery) (*[]swagger.AdResponse, error) {
	sqlStatement := `
	SELECT title, end_at FROM search_ads(
		:age, :gender, :country, :platform, :offset_val, :limit_val
	)
	`

	var adResponse []swagger.AdResponse
	rows, err := p.db.NamedQuery(sqlStatement, adQuery)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	for rows.Next() {
		var ad swagger.AdResponse
		if err = rows.StructScan(&ad); err != nil {
			logrus.Error(err)
			return nil, err
		}
		adResponse = append(adResponse, ad)
	}

	return &adResponse, nil
}
