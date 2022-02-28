package endpoint

import (
	"context"

	"github.com/khakim88/test-promo/model"
	"github.com/khakim88/test-promo/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ValidationPromotion endpoint.Endpoint
}

func MakeEndpoints(svc service.PromoService) Endpoints {
	return Endpoints{
		ValidationPromotion: MakeValidationPromotion(svc),
	}
}

func MakeValidationPromotion(svc service.PromoService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (response interface{}, err error) {
		requ := req.(model.ValidatePromotionRequest)
		return svc.ValidatePromotionService(ctx, &requ)
	}
}
