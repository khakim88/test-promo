package repository

import "github.com/khakim88/test-promo/model"

type DBReaderWriter interface {
	GetProductBySKU(sku string) (*model.Product, error)
}
