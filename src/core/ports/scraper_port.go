package ports

import "github.com/unawaretub86/club-hub/src/core/domain"

type ScrapperPort interface {
	ScrapCompanyData(franchises []domain.Franchise) ([]domain.FranchiseScrapData, error)
}
