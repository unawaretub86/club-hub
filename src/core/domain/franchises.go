package domain

type (
	Franchise struct {
		Name     *string  `json:"name,omitempty"`
		Url      *string  `json:"url,omitempty"`
		Location Location `json:"location,omitempty"`
	}

	Franchises []Franchise
)
