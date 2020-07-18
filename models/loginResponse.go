package models

// LoginResponse is where we store the token
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
