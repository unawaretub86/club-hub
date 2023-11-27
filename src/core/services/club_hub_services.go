package services

import (
	"github.com/unawaretub86/club-hub/src/core/domain"
	"github.com/unawaretub86/club-hub/src/core/ports"
)

type clubHubService struct {
	repository ports.RepositoryPort
	scraper    ports.ScrapperPort
}

func NewClubHubService(repository ports.RepositoryPort, scraper ports.ScrapperPort) *clubHubService {
	return &clubHubService{
		repository,
		scraper,
	}
}

func (s *clubHubService) SaveCompany(company domain.Company) (*domain.Company, error) {

	information, err := s.scraper.ScrapCompanyData(company.Franchises)
	if err != nil {
		return nil, err
	}

	company.FranchiseScrapData = information

	return s.repository.SaveCompany(company)
}

func (s *clubHubService) GetCompany(filterFields map[string]string) (domain.Companies, error) {
	return s.repository.GetCompany(filterFields)
}

func (s *clubHubService) UpdateCompany(id uint, company domain.Company) (*domain.Company, error) {
	return s.repository.UpdateCompany(id, company)
}

func (s *clubHubService) GetCompanyByFranchise(filterFields map[string]string) (*domain.Company, error) {
	return s.repository.GetCompanyByFranchise(filterFields)
}

func (s *clubHubService) GetCompanyByInformation(filterFields map[string]string) (*domain.Company, error) {
	return s.repository.GetCompanyByInformation(filterFields)
}

func (s *clubHubService) GetCompanyByOwner(filterFields map[string]string) (*domain.Company, error) {
	return s.repository.GetCompanyByOwner(filterFields)
}
