package buyers

import (
	"backend-restaurant-transactions-visualizer/internal/models"
	"log"
)

type Service interface {
	FindAllBuyers() ([]models.Buyer, error)
}

type buyerService struct {
	buyerRepo Repository
}

func NewBuyersService(buyerRepo Repository) *buyerService {
	return &buyerService{buyerRepo: buyerRepo}
}

func (s *buyerService) FindAllBuyers() ([]models.Buyer, error) {

	buyers, err := s.buyerRepo.FindAllBuyers()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return buyers, nil
}
