package repository

import (
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

func (repo *ClubRepository) Save(company domain.Company) (*domain.Company, error) {
	result := repo.db.Create(&company)
	if result.Error != nil {
		return nil, result.Error
	}

	return &company, nil
}

func (repo *ClubRepository) Get(filterFields map[string]string) (*domain.Company, error) {
	company := &domain.Company{}

	result := repo.db.Where(filterFields).Find(company)
	if result.Error != nil {
		return nil, result.Error
	}

	return company, nil
}

func (repo *ClubRepository) Update(id uint, company domain.Company) (*domain.Company, error) {
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
