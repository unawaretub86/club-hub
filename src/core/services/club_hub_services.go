package services

import (
	"github.com/unawaretub86/club-hub/src/core/domain"
	"github.com/unawaretub86/club-hub/src/core/ports"
)

type clubHubService struct {
	repository ports.RepositoryPort
}

func NewClubHubService(repository ports.RepositoryPort) *clubHubService {
	return &clubHubService{
		repository,
	}
}

func (s *clubHubService) SaveCompany(company domain.Company) (*domain.Company, error) {
	return s.repository.Save(company)
}

func (s *clubHubService) GetCompany(filterFields map[string]string) (*domain.Company, error) {
	return s.repository.Get(filterFields)
}

func (s *clubHubService) UpdateCompany(id uint, company domain.Company) (*domain.Company, error) {
	return s.repository.Update(id, company)
}
