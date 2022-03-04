package models

type Role struct {
	ApplicationID *string  `json:"application_id,omitempty"`
	Roles         []string `json:"roles,omitempty"`
}
