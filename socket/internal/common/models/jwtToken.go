package models

type JWTToken struct {
	Exp        int64      `json:"exp"`
	FamilyName string   `json:"family_name"`
	FullName   string   `json:"full_name"`
	GivenName  string   `json:"given_name"`
	HearingID  []string `json:"hearing_id"`
	Iat        int64    `json:"iat"`
	Iss        string   `json:"iss"`
	Nbf        int64      `json:"nbf"`
	SessionID  []string `json:"session_id"`
	UserEmail  string   `json:"user_email"`
	UserID     string   `json:"user_id"`
}
