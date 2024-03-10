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
	return au.adRepo.CreateAd(ctx, ad)
}
