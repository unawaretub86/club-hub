package domain

type (
	FranchiseScrapData struct {
		ID                uint   `gorm:"primaryKey" json:"id,omitempty"`
		ImageURL          string `json:"image_url,omitempty"`
		IsWebsiteUp       bool   `json:"is_website_up,omitempty"`
		CommunicationType string `json:"communication_type,omitempty"`
		HopCount          int    `json:"hop_count,omitempty"`
		Servers           string
		DomainScrapDataID uint `json:"domainScrapDataId,omitempty"`
		DomainScrapData   DomainScrapData
		CompanyID         uint `json:"-"`
	}

	DomainScrapData struct {
		ID           uint   `gorm:"primaryKey" json:"id,omitempty"`
		CreatedAt    string `json:"created_at,omitempty"`
		ExpiresAt    string `json:"expires_at,omitempty"`
		Registrant   string `json:"registrant,omitempty"`
		ContactEmail string `json:"contact_email,omitempty"`
		CompanyID    uint   `json:"-"`
	}
)
