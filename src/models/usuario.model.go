package models

type Appcomercial struct {
	Code          int64  `json:"code" gorm:"not null" validate:"nonzero"`
	Access_Token  string `json:"access_token" gorm:"not null" validate:"nonzero"`
	Client_id     string `json:"client_id" gorm:"not null" validate:"nonzero"`
	Client_secret string `json:"client_secret" gorm:"not null" validate:"nonzero"`
}

func ValidadadosUsuario() {

}
