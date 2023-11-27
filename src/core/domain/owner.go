package domain

type (
	Owner struct {
		ID        *uint    `gorm:"primaryKey" json:"id,omitempty"`
		FirstName *string  `gorm:"not null" json:"first_name,omitempty"`
		LastName  *string  `gorm:"not null" json:"last_name,omitempty"`
		ContactID *uint    `gorm:"references:contacts" json:"contactId,omitempty"`
		Contact   *Contact `json:"contact,omitempty"`
	}

	Contact struct {
		ID         *uint     `gorm:"primaryKey" json:"id,omitempty"`
		Email      *string   `gorm:"not null,unique" json:"email,omitempty"`
		Phone      *string   `gorm:"not null,unique" json:"phone,omitempty"`
		LocationID *uint     `gorm:"references:locations" json:"locationId,omitempty"`
		Location   *Location `json:"location,omitempty"`
	}
)
