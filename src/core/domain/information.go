package domain

type Information struct {
	ID         *uint     `gorm:"primaryKey" json:"id,omitempty"`
	Name       *string   `gorm:"not null" json:"name,omitempty"`
	TaxNumber  *string   `gorm:"not null" json:"tax_number,omitempty"`
	LocationID *uint     `gorm:"references:locations" json:"locationId,omitempty"`
	Location   *Location `json:"location,omitempty"`
}
