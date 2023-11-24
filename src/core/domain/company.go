package domain

type Company struct {
	ID          uint        `json:"id"`
	Owner       Owner       `json:"owner,omitempty"`
	Information Information `json:"information,omitempty"`
	Franchises  Franchises  `json:"franchises,omitempty"`
}
