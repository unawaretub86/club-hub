package ports

import "github.com/unawaretub86/club-hub/src/core/domain"

type WebPort interface {
	SaveCompany(domain.Company) (*domain.Company, error)
	GetCompany(map[string]string) (*domain.Company, error)
	UpdateCompany(uint, domain.Company) (*domain.Company, error)
}
