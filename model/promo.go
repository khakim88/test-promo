package model

type ValidatePromotionRequest struct {
	PromotionCode string        `json:"promotion_code"`
	ProductCart   []ProductItem `json:"product_cart"`
}

type ProductItem struct {
	SkuProduct string `json:"sku_product"`
	Quantity   int64  `json:"quantity"`
}

type ValidatePromotionResponse struct {
	DiscountType   string  `json:"discount_type"`
	DiscountAmount float64 `json:"discount_amount"`
	TotalPrice     float64 `json:"discount_amount"`
	Description    string  `json:"description"`
}
