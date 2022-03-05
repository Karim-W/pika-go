package models

type PresenceUpdate struct {
	AddedUsers   []string `json:"addedUsers"`
	RemovedUsers []string `json:"removedUsers"`
	SessionId    string   `json:"sessionId"`
}
