package model

type Product struct {
	SkuProduct   string  `json:"sku_product"`
	Quantity     int64   `json:"quantity"`
	Price        float64 `json:"price"`
	Name         string  `json:"name"`
	InventoryQty int64   `json:"inventory_qty"`
}
