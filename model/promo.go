package model

type ValidatePromotionRequest struct {
	PromotionCode string `json:"promotion_code"`
}

type ValidatePromotionResponse struct {
	Valid          bool    `json:"valid"`
	DiscountType   string  `json:"discount_type"`
	DiscountAmount float64 `json:"discount_amount"`
	Description    string  `json:"description"`
	MaxDiscount    float64 `json:"max_discount"`
}
