package models

import "encoding/json"

type Message struct {
  Audience []string `json:"audience"`
  AudienceType string  `json:"audienceType"`
  Command string `json:"command"`
  Data interface{} `json:"data"`
  RetryCount int `json:"retryCount,omitempty"`
}

func (m *Message) ConvertToBytes(msg Message) *[]byte {
  bytes, err := json.Marshal(msg)
  if err != nil {
    return nil
  }
  return &bytes
}

func ParseMessageFromBytes(bt *[]byte) *Message{
  msg := Message{}
  json.Unmarshal(*bt,&msg)
  return &msg
}
