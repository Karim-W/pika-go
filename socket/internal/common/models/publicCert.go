package models

type PublicCert struct{
  PublicKey string `json:"publicKey"`
  Kid string `json:"kid"`
}
