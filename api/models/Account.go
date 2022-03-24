package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	ID        uint32    `gorm:"primary_key; auto_increment" json:"id"`
	UserName  string    `gorm:"size:100;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"dafault:CURRENT_TIMESTAMP" json:"created_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassowrd, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassowrd), []byte(password))
}

//Before saving to DB, Hash Password return nil if no error
func (a *Account) BeforeSave() error {
	hashedPassword, err := Hash(a.Password)
	if err != nil {
		return err
	}

	a.Password = string(hashedPassword)
	return nil
}

func (a *Account) Prepare() {
	a.ID = 0
	a.UserName = html.EscapeString(strings.TrimSpace(a.UserName))
	a.CreatedAt = time.Now()
}

func (a *Account) Validate(action string) error {
	if a.UserName == "" {
		return errors.New("Require UserName")
	}
	if a.Password == "" {
		return errors.New("Require Passowrd")
	}
	return nil
}

func (a *Account) SaveAccount(db *gorm.DB) (*Account, error) {
	err := db.Debug().Create(&a).Error

	if err != nil {
		return &Account{}, err
	}
	return a, nil
}
