package domain

type Location struct {
	ID        *uint   `gorm:"primaryKey" json:"id,omitempty"`
	City      *string `gorm:"not null" json:"city,omitempty"`
	CountryID *uint   `gorm:"not null" json:"country_id,omitempty"`
	Address   *string `gorm:"not null" json:"address,omitempty"`
	ZipCode   *string `gorm:"not null" json:"zip_code,omitempty"`
}
