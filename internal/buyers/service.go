package buyers

import (
	"backend-restaurant-transactions-visualizer/internal/models"
	"log"
)

type Service interface {
	FindAllBuyers() (models.BuyerList, error)
	FindBuyerById(buyerId string) (*models.BuyerDetails, error)
}

type buyerService struct {
	buyerRepo Repository
}

func NewBuyersService(buyerRepo Repository) *buyerService {
	return &buyerService{buyerRepo: buyerRepo}
}

func (s *buyerService) FindAllBuyers() (models.BuyerList, error) {

	buyers, err := s.buyerRepo.FindAllBuyers()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return buyers, nil
}

func (s *buyerService) FindBuyerById(buyerId string) (*models.BuyerDetails, error) {

	buyerDetails, err := s.buyerRepo.FindBuyerById(buyerId)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return buyerDetails, nil
}
