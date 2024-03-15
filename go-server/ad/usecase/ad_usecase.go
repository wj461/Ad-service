package usecase

import (
	"context"

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
	if ad.Conditions.AgeEnd == 0 {
		ad.Conditions.AgeEnd = 100
	}

	return au.adRepo.CreateAd(ctx, ad)
}

func (au *adUsecase) SearchAd(ctx context.Context, adQuery *swagger.AdQuery) (*[]swagger.AdResponse, error) {
	return au.adRepo.SearchAd(ctx, adQuery)
}

func (au *adUsecase) ResetDB(ctx context.Context) error {
	return au.adRepo.ResetDB(ctx)
}
