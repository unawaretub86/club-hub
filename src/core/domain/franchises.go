package domain

type (
	Franchise struct {
		ID         *uint     `gorm:"primaryKey" json:"id,omitempty"`
		CompanyID  *uint     `gorm:"not null" json:"companyId,omitempty"`
		Name       *string   `gorm:"not null" json:"name,omitempty"`
		URL        *string   `gorm:"not null" json:"url,omitempty"`
		LocationID *uint     `gorm:"references:locations" json:"locationId,omitempty"`
		Location   *Location `json:"location,omitempty"`
	}

	Franchises []Franchise
)
