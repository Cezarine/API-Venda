package models

type Appscomercial struct {
	Code          int64  `json:"id"`
	Access_token  string `json:"access_token"`
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
}
