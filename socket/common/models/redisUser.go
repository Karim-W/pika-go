package models

type RedisUser struct {
  FirstName string `json:"firstName,omitempty"`
  LastName string  `json:"lastName,omitempty"`
  Email string `json:"email,omitempty"`
  ServerInstance string `json:"serverInstance,omitempty"`
  SocketId string `json:"socketId,omitempty"`
  Hearings []string `json:"hearings,omitempty"`
  Sessions []string `json:"sessions,omitempty"`
  State string `json:"state,omitempty"`
  UserType string `json:"userType,omitempty"`
  Id string `json:"id,omitempty"`
}
