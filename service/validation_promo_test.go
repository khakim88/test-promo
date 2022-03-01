package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/khakim88/test-promo/model"
	"github.com/khakim88/test-promo/repository"
)

func Test_promoService_ValidatePromotionService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockDBReaderWriter(ctrl)

	printCtx := context.Background()
	type args struct {
		ctx     context.Context
		request *model.ValidatePromotionRequest
	}
	tests := []struct {
		name     string
		fields1  *gomock.Call
		fields2  *gomock.Call
		fields3  *gomock.Call
		args     args
		wantResp *model.ValidatePromotionResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		//macbook pro & raspberry pi exp total 5.39
		{
			"test macbook & RaspBerry Pi",
			mockRepo.EXPECT().GetProductBySKU("43N23P").Return(&model.Product{
				SkuProduct: "43N23P",
				Name:       "MacBookPro",
				Price:      5399.99,
			}, nil),
			mockRepo.EXPECT().GetProductBySKU("234234").Return(&model.Product{
				SkuProduct: "234234",
				Name:       "RaspBerryPI",
				Price:      30.00,
			}, nil),
			nil,

			args{ctx: printCtx, request: &model.ValidatePromotionRequest{ProductCart: []model.ProductItem{
				{
					SkuProduct: "43N23P",
					Quantity:   1,
				}, {
					SkuProduct: "234234",
					Quantity:   1,
				},
			}}},
			&model.ValidatePromotionResponse{
				TotalPrice:     5399.99,
				DiscountAmount: 30.00,
			},
			true,
		},
		{
			"googl home > 3",
			mockRepo.EXPECT().GetProductBySKU("120P90").Return(&model.Product{
				SkuProduct: "120P90",
				Name:       "google home",
				Price:      49.99,
			}, nil),
			nil,
			nil,

			args{ctx: printCtx, request: &model.ValidatePromotionRequest{ProductCart: []model.ProductItem{
				{
					SkuProduct: "120P90",
					Quantity:   3,
				},
			}}},
			&model.ValidatePromotionResponse{
				TotalPrice:     99.98,
				DiscountAmount: 49.99,
			},
			true,
		},
		{
			"A304SD > 3 =10%",
			mockRepo.EXPECT().GetProductBySKU("A304SD").Return(&model.Product{
				SkuProduct: "A304SD",
				Name:       "Alexa Speaker",
				Price:      109.50,
			}, nil),
			nil,
			nil,

			args{ctx: printCtx, request: &model.ValidatePromotionRequest{ProductCart: []model.ProductItem{
				{
					SkuProduct: "A304SD",
					Quantity:   3,
				},
			}}},
			&model.ValidatePromotionResponse{
				TotalPrice:     295.65,
				DiscountAmount: 32.85,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &promoService{
				repo: mockRepo,
			}
			gotResp, err := rs.ValidatePromotionService(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("promoService.ValidatePromotionService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("promoService.ValidatePromotionService() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
