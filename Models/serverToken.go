package models

type ServerToken struct {
	Aud      []string `json:"aud"`
	ClientID string   `json:"client_Id"`
	Dat      string   `json:"dat"`
	Exp      int      `json:"exp"`
	Iat      int      `json:"iat"`
	Iss      string   `json:"iss"`
	Nbf      int      `json:"nbf"`
	Scope    []string `json:"scope"`
}
