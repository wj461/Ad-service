package domain

import "github.com/wj461/ad-service/swagger"

type adRepository interface {
	CreateAd(ad *swagger.Ad) error
}

type adUsecase interface {
	CreateAd(ad *swagger.Ad) error
}
