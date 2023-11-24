package domain

type Location struct {
	City    *string `json:"phone,omitempty"`
	Country *string `json:"Country,omitempty"`
	Address *string `json:"Address,omitempty"`
	ZipCode *string `json:"zip_code,omitempty"`
}
