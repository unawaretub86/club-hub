package domain

type Information struct {
	FirstName *string `json:"first_name,omitempty"`
	TaxNumber *string `json:"tax_number,omitempty"`
	Location Location `json:"location,omitempty"`
}
