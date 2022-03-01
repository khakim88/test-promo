package service

import (
	"context"
	"math"

	"github.com/khakim88/test-promo/model"
)

func (rs *promoService) ValidatePromotionService(ctx context.Context, request *model.ValidatePromotionRequest) (resp *model.ValidatePromotionResponse, err error) {

	//DEFAULT rule
	freeRaspBerry := 0
	var discountAmount, priceAmount float64
	for _, prod := range request.ProductCart {
		if prod.SkuProduct == "43N23P" {

			prodMac, _ := rs.repo.GetProductBySKU(prod.SkuProduct)

			p := float64(prod.Quantity) * prodMac.Price
			priceAmount = priceAmount + p
			//get free
			freeRaspBerry = int(prod.Quantity)

		}
		//each macbookpro free raspberry
		if prod.SkuProduct == "234234" {
			if freeRaspBerry > 0 && prod.Quantity > 0 {
				prodRaspBerry, _ := rs.repo.GetProductBySKU(prod.SkuProduct)
				//discount price raspberry
				priceAmount = priceAmount + (prodRaspBerry.Price * float64(prodRaspBerry.Quantity))
				discountAmount = discountAmount + (prodRaspBerry.Price * float64(prod.Quantity))
			}

		}
		//buying 3 google pay 2 price / free 1
		if prod.SkuProduct == "120P90" {
			if prod.Quantity > 3 {
				prodGoogle, _ := rs.repo.GetProductBySKU(prod.SkuProduct)
				calc := float64(prod.Quantity) / 3
				disc := math.Floor(calc) * prodGoogle.Price
				priceAmount = priceAmount + (prodGoogle.Price * float64(prodGoogle.Quantity))
				discountAmount = discountAmount + disc

			}
		}
		//buying more than 3 alexa
		if prod.SkuProduct == "A304SD" {
			if prod.Quantity > 3 {
				prodAlexa, _ := rs.repo.GetProductBySKU(prod.SkuProduct)
				price := prodAlexa.Price * float64(prodAlexa.Quantity)
				disc := price * 0.1
				priceAmount = priceAmount + (prodAlexa.Price * float64(prodAlexa.Quantity))
				discountAmount = discountAmount + disc

			}
		}

	}
	resp.DiscountAmount = discountAmount
	resp.TotalPrice = priceAmount - discountAmount

	return resp, nil
}
