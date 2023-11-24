package domain

type Company struct {
	Owner       Owner       `json:"owner,omitempty"`
	Information Information `json:"information,omitempty"`
	Franchises  Franchises  `json:"franchises,omitempty"`
}
