package product

import (
	"github.com/google/uuid"
	"github.com/suadev/go-rest-api-clean-architecture/entity"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateProduct(product *entity.Product) (*entity.Product, error) {
	product.ID = uuid.New()
	createdProduct, err := s.repo.Create(product)

	if err != nil {
		return createdProduct, err
	}
	return createdProduct, nil
}

func (s *Service) BulkUpdate(products *[]entity.Product) error {
	return s.repo.BulkUpdate(products)
}

func (s *Service) GetProducts() ([]entity.Product, error) {
	return s.repo.GetList()
}

func (s *Service) GetProduct(id uuid.UUID) (entity.Product, error) {
	return s.repo.GetById(id)
}
