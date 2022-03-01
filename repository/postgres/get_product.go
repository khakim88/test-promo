package postgres

import (
	"github.com/khakim88/test-promo/common/logger"
	"github.com/khakim88/test-promo/model"
)

func (p *postgresConn) GetProductBySKU(sku string) (*model.Product, error) {
	product := new(model.Product)
	query := `SELECT 
				sku,
				name,
				price,
				qty
			FROM product
			WHERE sku = $1`

	err := p.db.QueryRow(query, sku).Scan(
		&product.SkuProduct,
		&product.Name,
		&product.Price,
		&product.Quantity,
	)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return product, nil
}
