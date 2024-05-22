package model

type UserGoogle struct {
	ID         string `json:"id,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	Email      string `json:"email,omitempty"`
}
