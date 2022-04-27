package models

import (
	"fmt"

	"gorm.io/gorm"
)

// this struct allows a user to submit a turn
type TurnInput struct {
	Username string `json:"username"`
	GameName string `json:"gamename"`
	TileNum  int    `json:"tilenum"`
}

// this struct represents a board of tic tac toe
type GameBoard struct {
	gorm.Model
	BoardName string `json:"boardname"`
	SqOne     string `json:"sqone"`
	SqTwo     string `json:"sqtwo"`
	SqThree   string `json:"sqthree"`
	SqFour    string `json:"sqfour"`
	SqFive    string `json:"sqfive"`
	SqSix     string `json:"sqsix"`
	SqSeven   string `json:"sqseven"`
	SqEight   string `json:"sqeight"`
	SqNine    string `json:"sqnine"`
}

// this struct encapsulates all data needed for a game
type Game struct {
	gorm.Model
	GameName      string `json:"gamename"`
	BoardID       int
	Board         GameBoard `gorm:"foreignKey:BoardID"`
	PlOneID       int
	PlOne         TicUser `gorm:"foreignKey:PlOneID"`
	PlTwoID       int
	PlTwo         TicUser `gorm:"foreignKey:PlTwoID"`
	IsWinner      bool    `json:"iswinner"`
	CurrentPlayer string  `json:"currentplayer"`
	WinningPlayer string  `json:"winningplayer"`
}

// this struct takes in info to create a new game
type NewGameInfo struct {
	GameName string `json:"gamename"`
	PlOne    string `json:"plone"`
	PlTwo    string `json:"pltwo"`
}

// this struct takes in info to delete a game
type DeleteGame struct {
	GameName string `json:"gamename"`
}

// this struct takes in info to get a games status
type GameStatus struct {
	GameName string `json:"gamename"`
}

// this method checks if a user has won a game and modifies the game object
func (G *Game) Checkforwin(B *GameBoard) bool {
	fmt.Println("did you win?")
	return true
}

// this method returns the list of users that are in a game(only two)
func (G *Game) ListUsers() []TicUser {
	gameList := []TicUser{G.PlOne, G.PlTwo}
	return gameList
}

// this method switches the players whose turn it is and takes in [username]
func (G *Game) SwitchPlayer(name string) {
	fmt.Println("Players Switched")
}

// this method sets the current player
func (G *Game) SetCurrentPlayer(newUsername string) {
	G.CurrentPlayer = newUsername
}
