package domain

type (
	Company struct {
		ID            *uint        `gorm:"primaryKey" json:"id,omitempty"`
		OwnerID       *uint        `gorm:"references:owners" json:"ownerId,omitempty"`
		Owner         *Owner       `json:"owner,omitempty"`
		InformationID *uint        `gorm:"references:informacion" json:"informacionId,omitempty"`
		Information   *Information `json:"informacion,omitempty"`
		Franchises    *Franchises  `gorm:"foreignKey:CompanyID" json:"franchises,omitempty"`
	}
	ReqData struct {
		Company Company `json:"company"`
	}

	Companies []Company
)
