package usecases

import (
	"github.com/unawaretub86/club-hub/src/core/domain"
	"github.com/unawaretub86/club-hub/src/core/ports"
)

type ScrapperUseCase struct {
	rest ports.ScrapperPort
}

func NewScrapperUseCase(rest ports.ScrapperPort) *ScrapperUseCase {
	return &ScrapperUseCase{
		rest,
	}
}

func (UseCase *ScrapperUseCase) ScrapCompanyData(franchises []domain.Franchise) ([]domain.FranchiseScrapData, error) {
	return UseCase.rest.ScrapCompanyData(franchises)
}
