package domain

import "github.com/lib/pq"

type (
	FranchiseScrapData struct {
		ID                uint           `gorm:"primaryKey" json:"id,omitempty"`
		ImageURL          string         `json:"image_url,omitempty"`
		Status       string         `json:"status,omitempty"`
		CommunicationType string         `json:"communication_type,omitempty"`
		HopCount          int            `json:"hop_count,omitempty"`
		Servers           pq.StringArray `gorm:"type:text[]" json:"servers,omitempty"`
		DomainScrapDataID uint           `json:"domainScrapDataId,omitempty"`
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

	SSLDomainScrapData struct {
		Host     string    `json:"host"`
		Port     int       `json:"port"`
		Protocol string    `json:"protocol"`
		Status   string    `json:"status"`
		Servers  []Servers `json:"endpoints"`
	}

	Servers struct {
		ServerName string `json:"serverName"`
	}
)
