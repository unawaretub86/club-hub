package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/unawaretub86/club-hub/src/core/domain"
)

type ClubRepository struct {
	db *gorm.DB
}

func NewClubRepository(db *gorm.DB) *ClubRepository {
	return &ClubRepository{
		db: db,
	}
}

func (repo *ClubRepository) SaveCompany(company domain.Company) (*domain.Company, error) {

	fmt.Println(company)

	if err := repo.db.Create(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (repo *ClubRepository) GetCompany(filterFields map[string]string) (*domain.Company, error) {
	company := &domain.Company{}

	result := repo.db.Where(filterFields).Find(company)
	if result.Error != nil {
		return nil, result.Error
	}

	return company, nil
}

func (repo *ClubRepository) UpdateCompany(id uint, company domain.Company) (*domain.Company, error) {
	companyResult, result := repo.getByID(id)
	if result.Error != nil {
		return nil, result.Error
	}

	company.ID = companyResult.ID

	result = repo.db.Updates(company)
	if result.Error != nil {
		return nil, result.Error
	}

	return &company, nil
}

func (repo *ClubRepository) getByID(id uint) (*domain.Company, *gorm.DB) {
	company := &domain.Company{}

	result := repo.db.Take(&company, id)

	return company, result
}
