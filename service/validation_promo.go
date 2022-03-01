package service

import (
	"context"
	"math"

	"github.com/khakim88/test-promo/common/logger"
	"github.com/khakim88/test-promo/model"
)

func (rs *promoService) ValidatePromotionService(ctx context.Context, request *model.ValidatePromotionRequest) (resp *model.ValidatePromotionResponse, err error) {

	//DEFAULT rule
	response := new(model.ValidatePromotionResponse)
	freeRaspBerry := 0
	var discountAmount, priceAmount, TotalPrice float64
	for _, prod := range request.ProductCart {
		if prod.SkuProduct == "43N23P" {

			prodMac, _ := rs.repo.GetProductBySKU(prod.SkuProduct)

			p := float64(prod.Quantity) * prodMac.Price
			priceAmount = priceAmount + p
			//get free
			freeRaspBerry = int(prod.Quantity)
			TotalPrice = priceAmount - discountAmount

		}
		//each macbookpro free raspberry
		if prod.SkuProduct == "234234" {
			if freeRaspBerry > 0 && prod.Quantity > 0 {
				prodRaspBerry, _ := rs.repo.GetProductBySKU(prod.SkuProduct)
				//discount price raspberry
				priceAmount = priceAmount + (prodRaspBerry.Price * float64(prod.Quantity))
				discountAmount = discountAmount + (prodRaspBerry.Price * float64(prod.Quantity))
				TotalPrice = priceAmount - discountAmount
			}

		}
		//buying 3 google pay 2 price / free 1
		if prod.SkuProduct == "120P90" {
			if prod.Quantity >= 3 {
				prodGoogle, _ := rs.repo.GetProductBySKU(prod.SkuProduct)
				calc := float64(prod.Quantity) / 3

				c := math.Floor(calc)
				disc := c * prodGoogle.Price
				priceAmount = priceAmount + (prodGoogle.Price * float64(prod.Quantity))
				discountAmount = (discountAmount + disc)
				TotalPrice = priceAmount - discountAmount
				TotalPrice = (math.Ceil(TotalPrice*100) / 100)
				logger.Info("[CEIL]:", math.Floor(calc), "[PRICEAMOUNT]:", priceAmount, "[disc]:", disc, "[discountAmount]:", discountAmount, "[TotalPrice]:", TotalPrice)

			}
		}
		//buying more than 3 alexa
		if prod.SkuProduct == "A304SD" {
			if prod.Quantity >= 3 {
				prodAlexa, _ := rs.repo.GetProductBySKU(prod.SkuProduct)
				price := prodAlexa.Price * float64(prod.Quantity)
				disc := price * 0.1
				priceAmount = priceAmount + price
				discountAmount = (discountAmount + disc)
				TotalPrice = priceAmount - discountAmount

			}
		}

	}

	response.DiscountAmount = discountAmount
	response.TotalPrice = TotalPrice

	return response, nil
}
