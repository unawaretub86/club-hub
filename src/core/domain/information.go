package domain

type Information struct {
	Name *string `json:"name,omitempty"`
	TaxNumber *string `json:"tax_number,omitempty"`
	Location Location `json:"location,omitempty"`
}
