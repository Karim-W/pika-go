package models

type SocketPayload struct {
	Audience     []string          `json:"audience"`
	AudienceType string            `json:"audienceType"`
	Command      string            `json:"command"`
	Data         map[string]string `json:"data"`
}
