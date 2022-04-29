package main

import (
	"TicTac_Api/database"
	"TicTac_Api/models"
	"TicTac_Api/routes"
	"fmt"

	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// function that creates all routes
func initialize_routes(app *fiber.App) {
	// routes that interact with users
	app.Get("/ticgame/users/all", routes.GetAllUsers)        // get a slice of all users(Tested)
	app.Post("/ticgame/users/add", routes.AddUser)           // add a new users(Tested)
	app.Post("/ticgame/users/user", routes.GetUser)          // get a single user(Tested)
	app.Post("/ticgame/users/update", routes.ChangePassword) // change user password(Tested)
	app.Post("/ticgame/users/delete", routes.DeleteUser)     // delete a user(Tested)

	// routes that interact with games
	app.Post("/ticgame/games/new", routes.NewGameStart) // start a new game(Tested)
	app.Post("/ticgame/delete", routes.EndGame)         // delete a game and corresponding game board(Tested)
	app.Post("/ticgame/games/turn", routes.TakeTurn)    // take a turn for a player(Tested)
	app.Get("/ticgame/games/all", routes.GetAllGames)   // return a json list of all games and their info(Tested)

	// routes that interact with boards
	app.Post("/ticgame/boards/status", routes.BoardState) // get a boards status(Tested)
	app.Get("/ticgame/boards/all", routes.GetAllBoards)   // get all boards statuses(Tested)
}

// function that opens db connection and migrates tables
func initialize_database() {
	var err error
	database.DBC, err = gorm.Open(sqlite.Open("TicTac.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	fmt.Println("Connection to database now open")
	database.DBC.AutoMigrate(&models.Game{}, &models.TicUser{}, &models.GameBoard{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New() // create new fiber instance

	initialize_database() // create or migrate database

	initialize_routes(app) // create all routes

	log.Fatal(app.Listen(":3000")) // lsiten to a specific port
}
