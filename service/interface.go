package service

import (
	"context"

	"github.com/khakim88/test-promo/model"
)

type PromoService interface {
	ValidatePromotionService(ctx context.Context, request *model.ValidatePromotionRequest) (resp *model.ValidatePromotionResponse, err error)
}
