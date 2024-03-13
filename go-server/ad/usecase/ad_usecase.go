package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wj461/ad-service/domain"
	"github.com/wj461/ad-service/swagger"
)

type adUsecase struct {
	adRepo domain.AdRepository
}

func NewAdUsecase(adRepo domain.AdRepository) domain.AdUsecase {
	return &adUsecase{
		adRepo: adRepo,
	}
}

func (au *adUsecase) CreateAd(ctx context.Context, ad *swagger.Ad) error {
	if !ad.Conditions.Male && !ad.Conditions.Female {
		ad.Conditions.Male = true
		ad.Conditions.Female = true
	}

	if ad.Conditions.AgeEnd == 0 {
		ad.Conditions.AgeEnd = 999
	}

	return au.adRepo.CreateAd(ctx, ad)
}

func (au *adUsecase) SearchAd(ctx context.Context, adQuery *swagger.AdQuery) (*[]swagger.AdResponse, error) {
	logrus.Info(adQuery)
	return au.adRepo.SearchAd(ctx, adQuery)
}
