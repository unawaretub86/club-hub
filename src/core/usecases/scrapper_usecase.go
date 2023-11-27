package usecases

import (
	"fmt"

	"github.com/unawaretub86/club-hub/src/core/domain"
	"github.com/unawaretub86/club-hub/src/core/ports"
)

type ScrapperUseCase struct{
	rest ports.ScrapperPort
}

func NewScrapperUseCase(rest ports.ScrapperPort) *ScrapperUseCase {
	return &ScrapperUseCase{
		rest,
	}
}

func (UseCase *ScrapperUseCase) ScrapCompanyData(franchises []domain.Franchise) ([]domain.Franchise, error) {

	info, err := UseCase.rest.ScrapCompanyData(franchises)
	if err != nil {
		return nil, err
	}

	fmt.Println(info)

	return nil, nil
}
