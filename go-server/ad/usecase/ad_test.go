package usecase

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wj461/ad-service/swagger"

	adRepository "github.com/wj461/ad-service/ad/repository"
)

func init() {
	viper.SetConfigName(".env")
	viper.AddConfigPath("../../../")
	viper.SetConfigType("dotenv")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Error reading config file, ", err)
	}
}

func ConnectDB() *sqlx.DB {
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("POSTGRES_USER")
	dbPassword := viper.GetString("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	// if go not run in docker, host will be localhost
	if dbHost == "" {
		dbHost = "localhost"
	}

	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	return db
}

// func TestInsert(t *testing.T) {
// 	adPostgresqlRepository := adRepository.NewPostgresqlAdRepository(ConnectDB())
// 	var ad = &swagger.Ad{}
// 	ans := &swagger.AdData{}

// 	err := adPostgresqlRepository.CreateAd(nil, ad)
// 	if err != nil {
// 		t.Errorf("Error %s", err)
// 	}

// 	result, err := adPostgresqlRepository.CheckAd(nil, ad)
// 	if err != nil {
// 		t.Errorf("Error %s", err)
// 	}
// 	if result != ans {
// 		ansJson, _ := json.Marshal(ans)
// 		t.Errorf("Result was incorrect, got: %s, want: %s.", result, string(ansJson))
// 	}

// }

func TestSearch(t *testing.T) {
	adPostgresqlRepository := adRepository.NewPostgresqlAdRepository(ConnectDB())
	var adQuery = &swagger.AdQuery{
		OffsetVal: new(int),
		LimitVal:  new(int),
		Gender:    new(string),
		Country:   new(string),
		Platform:  new(string),
	}
	*adQuery.OffsetVal = 0
	*adQuery.LimitVal = 10
	*adQuery.Gender = "F"
	*adQuery.Country = "TW"
	*adQuery.Platform = "ios"

	ans := &[]swagger.AdResponse{
		{
			Title: "Ad 2",
			EndAt: time.Date(2024, time.August, 11, 12, 0, 0, 0, time.UTC),
		}, {
			Title: "Ad 3",
			EndAt: time.Date(2024, time.August, 12, 12, 0, 0, 0, time.UTC),
		}, {
			Title: "Ad 8",
			EndAt: time.Date(2024, time.August, 17, 12, 0, 0, 0, time.UTC),
		}, {
			Title: "Ad 10",
			EndAt: time.Date(2024, time.August, 19, 12, 0, 0, 0, time.UTC),
		},
	}
	result, err := adPostgresqlRepository.SearchAd(nil, adQuery)
	if err != nil {
		t.Errorf("Error %s", err)
	}

	if !compareAdResponses(*result, *ans) {
		t.Errorf("Result was incorrect)")
	}
}

func compareAdResponses(a, b []swagger.AdResponse) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Title != b[i].Title || !v.EndAt.Equal(b[i].EndAt) {
			return false
		}
	}
	return true
}
