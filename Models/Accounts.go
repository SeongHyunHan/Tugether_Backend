package models

type Accounts struct {
	AutoId   int    `json:"autoid" gorm:"primaryKey"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
