package repository

import (
	"reflect"

	"gorm.io/gorm"

	"github.com/unawaretub86/club-hub/src/config/apperrors"
	"github.com/unawaretub86/club-hub/src/core/domain"
)

const (
	id            = "ID"
	informationID = "information_id"
	ownerID       = "owner_id"
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
	if err := repo.db.Create(&company).Error; err != nil {
		return nil, err
	}

	return &company, nil
}

func (repo *ClubRepository) GetCompany(filterFields map[string]string) (domain.Companies, error) {
	company := domain.Companies{}

	result := repo.db.
		Preload("Owner.Contact.Location.Country").
		Preload("Information.Location.Country").
		Preload("Franchises.Location.Country").
		Where(filterFields).
		Find(&company)

	if result.Error != nil {
		return nil, result.Error
	}

	return company, nil
}

func (repo *ClubRepository) UpdateCompany(id uint, company domain.Company) (*domain.Company, error) {
	// Start a transaction
	tx := repo.db.Begin()

	// Update the main table (Company) using Updates
	if err := tx.Model(&company).Updates(&company).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update the associated tables (Owner, Information) using subqueries
	if company.Owner != nil {
		if err := tx.
			Model(&domain.Owner{}).
			Where("id = ?", &company.Owner.ID).
			Updates(&company.Owner).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if company.Information != nil {
		if err := tx.
			Model(&domain.Information{}).
			Where("id = ?", &company.Information.ID).
			Updates(&company.Information).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if company.Franchises != nil {
		for _, franchise := range company.Franchises {
			if err := tx.
				Model(&domain.Franchise{}).
				Where("id = ?", franchise.ID).
				Updates(&franchise).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	// Commit the transaction
	tx.Commit()

	return &company, nil
}

func (repo *ClubRepository) GetCompanyByFranchise(filterFields map[string]string) (*domain.Company, error) {
	franchises := []domain.Franchise{}

	result := repo.db.Where(filterFields).Find(&franchises)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(franchises) == 0 {
		return nil, apperrors.FranchiseNotFound
	}

	companyID := franchises[0].CompanyID

	company, err := repo.getCompanyBy(*companyID, "id")
	if err.Error != nil {
		return nil, err.Error
	}

	return company, nil
}

func (repo *ClubRepository) GetCompanyByInformation(filterFields map[string]string) (*domain.Company, error) {
	information := &domain.Information{}
	return repo.getCompanyByModel(filterFields, information, id, informationID)
}

func (repo *ClubRepository) GetCompanyByOwner(filterFields map[string]string) (*domain.Company, error) {
	owner := &domain.Owner{}
	return repo.getCompanyByModel(filterFields, owner, id, ownerID)
}

func (repo *ClubRepository) getCompanyBy(value uint, field string) (*domain.Company, *gorm.DB) {
	company := &domain.Company{}

	result := repo.db.
		Preload("Owner.Contact.Location").
		Preload("Information.Location").
		Preload("Franchises.Location").
		First(&company, field+" = ?", value)

	return company, result
}

func (repo *ClubRepository) getCompanyByModel(filterFields map[string]string, model interface{}, field, fieldToFind string) (*domain.Company, error) {
	result := repo.db.Where(filterFields).Find(model)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, apperrors.RecordNotFound
	}

	modelValue := reflect.ValueOf(model).Elem()
	modelID := modelValue.FieldByName(field).Elem().Uint()

	company, err := repo.getCompanyBy(uint(modelID), fieldToFind)
	if err.Error != nil {
		return nil, err.Error
	}

	return company, nil
}
