package domain

type (
	Owner struct {
		FirstName *string `json:"first_name,omitempty"`
		LastName  *string `json:"last_name,omitempty"`
		Contact   contact `json:"contact,omitempty"`
	}

	contact struct {
		Email    *string  `json:"email,omitempty"`
		Phone    *string  `json:"phone,omitempty"`
		Location Location `json:"location,omitempty"`
	}
)
