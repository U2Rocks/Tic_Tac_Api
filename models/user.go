package models

import (
	"gorm.io/gorm"
)

// this struct represents a basic user of the application
type TicUser struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Symbol   string `json:"symbol"` // text used to fill tic tac toe squares
}

// this struct takes in required data to grab a user
type GetTicUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// this struct takes in required data to change a password
type ChangeUserInfo struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

// this method changes a users symbol
func (T *TicUser) SetSymbol(newSymbol string) {
	T.Symbol = newSymbol
}
