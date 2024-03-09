package domain

type adRepository interface {
	CreateAd(ad *Ad) error
}

type adUsecase interface {
	CreateAd(ad *Ad) error
}
