package domain

type (
	Franchise struct {
		ID         *uint     `gorm:"primaryKey,unique" json:"id,omitempty"`
		CompanyID  *uint     `gorm:"not null,unique" json:"companyId,omitempty"`
		Name       *string   `gorm:"not null,unique" json:"name,omitempty"`
		URL        *string   `gorm:"not null,unique" json:"url,omitempty"`
		LocationID *uint     `gorm:"references:locations,unique" json:"locationId,omitempty"`
		Location   *Location `json:"location,omitempty"`
	}

	Franchises []Franchise
)
