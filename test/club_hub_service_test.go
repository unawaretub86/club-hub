package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/unawaretub86/club-hub/src/core/domain"
	"github.com/unawaretub86/club-hub/src/core/services"
	mock_ports "github.com/unawaretub86/club-hub/test/mocks"
)

type mocks struct {
	repository *mock_ports.MockRepositoryPort
	scraper    *mock_ports.MockScrapperPort
}

func TestGetCompany(t *testing.T) {
	m := mocks{
		repository: mock_ports.NewMockRepositoryPort(gomock.NewController(t)),
		scraper:    mock_ports.NewMockScrapperPort(gomock.NewController(t)),
	}

	mockClubHubService := services.NewClubHubService(m.repository, m.scraper)

	expectedFilterFields := map[string]string{}
	mockCompanies := domain.Companies{}
	m.repository.EXPECT().GetCompany(expectedFilterFields).Return(mockCompanies, nil)

	resultCompanies, err := mockClubHubService.GetCompany(expectedFilterFields)

	require.NoError(t, err)
	require.Len(t, resultCompanies, len(mockCompanies))
}

func TestUpdateCompany(t *testing.T) {
	m := mocks{
		repository: mock_ports.NewMockRepositoryPort(gomock.NewController(t)),
		scraper:    mock_ports.NewMockScrapperPort(gomock.NewController(t)),
	}

	mockClubHubService := services.NewClubHubService(m.repository, m.scraper)

	var id uint = 1
	name := "pablo"
	lastName := "chachon"

	expectedID := uint(1)
	expectedCompany := domain.Company{
		ID: &id,
		Owner: &domain.Owner{
			FirstName: &name,
			LastName:  &lastName,
		},
	}

	m.repository.EXPECT().UpdateCompany(expectedID, expectedCompany).Return(&expectedCompany, nil)

	resultCompany, err := mockClubHubService.UpdateCompany(expectedID, expectedCompany)

	assert.NoError(t, err)
	assert.NotNil(t, resultCompany)
	assert.Equal(t, expectedCompany, *resultCompany)
}

func TestGetCompanyByFranchise(t *testing.T) {
	m := mocks{
		repository: mock_ports.NewMockRepositoryPort(gomock.NewController(t)),
		scraper:    mock_ports.NewMockScrapperPort(gomock.NewController(t)),
	}

	mockClubHubService := services.NewClubHubService(m.repository, m.scraper)

	expectedFilterFields := map[string]string{"name": "Park royal"}

	var id uint = 1
	name := "Park royal"

	franchise := domain.Franchise{
		Name: &name,
	}

	franchises := []domain.Franchise{
		franchise,
	}

	expectedCompany := &domain.Company{
		ID:         &id,
		Franchises: franchises,
	}

	m.repository.EXPECT().GetCompanyByFranchise(expectedFilterFields).Return(expectedCompany, nil)

	resultCompany, err := mockClubHubService.GetCompanyByFranchise(expectedFilterFields)

	assert.NoError(t, err)
	assert.NotNil(t, resultCompany)
	assert.Equal(t, expectedCompany, resultCompany)
}

func TestGetCompanyByInformation(t *testing.T) {
	m := mocks{
		repository: mock_ports.NewMockRepositoryPort(gomock.NewController(t)),
		scraper:    mock_ports.NewMockScrapperPort(gomock.NewController(t)),
	}

	mockClubHubService := services.NewClubHubService(m.repository, m.scraper)

	expectedFilterFields := map[string]string{"name": "My entreprise holding"}

	var id uint = 1
	name := "Park royal"
	expectedCompany := &domain.Company{
		ID: &id,
		Information: &domain.Information{
			Name: &name,
		},
	}

	m.repository.EXPECT().GetCompanyByInformation(expectedFilterFields).Return(expectedCompany, nil)

	resultCompany, err := mockClubHubService.GetCompanyByInformation(expectedFilterFields)

	assert.NoError(t, err)
	assert.NotNil(t, resultCompany)
	assert.Equal(t, expectedCompany, resultCompany)
}

func TestGetCompanyByOwner(t *testing.T) {
	m := mocks{
		repository: mock_ports.NewMockRepositoryPort(gomock.NewController(t)),
		scraper:    mock_ports.NewMockScrapperPort(gomock.NewController(t)),
	}

	mockClubHubService := services.NewClubHubService(m.repository, m.scraper)

	expectedFilterFields := map[string]string{"first_name": "pablo"}

	var id uint = 1
	name := "pablo"
	expectedCompany := &domain.Company{
		ID: &id,
		Owner: &domain.Owner{
			FirstName: &name,
		},
	}

	m.repository.EXPECT().GetCompanyByOwner(expectedFilterFields).Return(expectedCompany, nil)

	resultCompany, err := mockClubHubService.GetCompanyByOwner(expectedFilterFields)

	assert.NoError(t, err)
	assert.NotNil(t, resultCompany)
	assert.Equal(t, expectedCompany, resultCompany)
}

func TestSaveCompany(t *testing.T) {
	m := mocks{
		repository: mock_ports.NewMockRepositoryPort(gomock.NewController(t)),
		scraper:    mock_ports.NewMockScrapperPort(gomock.NewController(t)),
	}

	mockClubHubService := services.NewClubHubService(m.repository, m.scraper)

	var id uint = 1
	var id2 uint = 2
	franchiseName1 := "marllot"
	franchiseName2 := "royal"
	name := "pablo"
	nameInformation := "My entreprise holding"

	expectedScrapData := []domain.FranchiseScrapData{
		{
			ID:                1,
			ImageURL:          "https://cdn2.paraty.es/parkroyal-corpo/images/9c795fe73b46a8a",
			Status:            "READY",
			CommunicationType: "http",
			HopCount:          2,
			DomainScrapDataID: 1,
			DomainScrapData: domain.DomainScrapData{
				ID:           1,
				CreatedAt:    "2006-09-28T17:37:03Z",
				ExpiresAt:    "2024-09-28T17:37:03",
				Registrant:   "Domain Privacy Service FBO Registrant.",
				ContactEmail: "park-royalhotels.com@domainprivacygroup.com",
			},
		},
		{
			ID:                2,
			ImageURL:          "https://cdn2.paraty.es/parkroyal-corpo/images/9c795fe73b46a8a",
			Status:            "READY",
			CommunicationType: "http",
			HopCount:          2,
			DomainScrapDataID: 1,
			DomainScrapData: domain.DomainScrapData{
				ID:           2,
				CreatedAt:    "2006-09-28T17:37:03Z",
				ExpiresAt:    "2024-09-28T17:37:03",
				Registrant:   "Domain Privacy Service FBO Registrant.",
				ContactEmail: "park-royalhotels.com@domainprivacygroup.com",
			},
		},
	}

	expectedCompany := domain.Company{
		ID:         &id,
		Owner: &domain.Owner{
			FirstName: &name,
		},
		Information: &domain.Information{
			Name: &nameInformation,
		},
		Franchises: []domain.Franchise{{ID: &id, Name: &franchiseName1}, {ID: &id2, Name: &franchiseName2}},
		FranchiseScrapData: expectedScrapData,
	}

	m.scraper.EXPECT().ScrapCompanyData(expectedCompany.Franchises).Return(expectedScrapData, nil)

	m.repository.EXPECT().SaveCompany(expectedCompany).Return(&expectedCompany, nil)

	resultCompany, err := mockClubHubService.SaveCompany(expectedCompany)

	assert.NoError(t, err)
	assert.NotNil(t, resultCompany)
	assert.Equal(t, &expectedCompany, resultCompany)
}
