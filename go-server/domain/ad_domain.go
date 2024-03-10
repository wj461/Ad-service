package domain

import (
	"context"

	"github.com/wj461/ad-service/swagger"
)

type AdRepository interface {
	CreateAd(ctx context.Context, ad *swagger.Ad) error
}

type AdUsecase interface {
	CreateAd(ctx context.Context, ad *swagger.Ad) error
}
