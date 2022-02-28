package service

import "github.com/khakim88/test-promo/repository"

type promoService struct {
	repo repository.DBReaderWriter
}

func NewPromoService(repo repository.DBReaderWriter) promoService {
	return promoService{
		repo: repo,
	}
}
