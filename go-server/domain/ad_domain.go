package domain

import (
	"context"

	"github.com/wj461/ad-service/swagger"
)

type AdRepository interface {
	CreateAd(ctx context.Context, ad *swagger.Ad) error
	// CheckAd(ctx context.Context, ad *swagger.Ad) (*swagger.AdData, error)
	SearchAd(ctx context.Context, adQuery *swagger.AdQuery) (*[]swagger.AdResponse, error)
}

type AdUsecase interface {
	CreateAd(ctx context.Context, ad *swagger.Ad) error
	SearchAd(ctx context.Context, adQuery *swagger.AdQuery) (*[]swagger.AdResponse, error)
}
