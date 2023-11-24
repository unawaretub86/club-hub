package ports

import "github.com/unawaretub86/club-hub/src/core/domain"

type RepositoryPort interface {
	Save(domain.Company) (*domain.Company, error)
	Get(map[string]string) (*domain.Company, error)
	Update(uint, domain.Company) (*domain.Company, error)
}
