package models

type UserToken struct {
	HearingID string `json:"HearingId,omitempty"`
	Issuer    string `json:"Issuer,omitempty"`
	Roles     []Role `json:"Roles,omitempty"`
	SessionID string `json:"SessionId,omitempty"`
	UserEmail string `json:"UserEmail,omitempty"`
	UserID    string `json:"UserId,omitempty"`
	Exp       int64  `json:"exp,omitempty"`
	Iat       int64  `json:"iat,omitempty"`
	Iss       string `json:"iss,omitempty"`
	Nbf       int64  `json:"nbf,omitempty"`
	Type      string `json:"type,omitempty"`
}
