package routes

import (
	"TicTac_Api/helperfunctions"

	"TicTac_Api/models"

	"TicTac_Api/database"

	"github.com/gofiber/fiber/v2"
)

// this function creates a new game
func NewGameStart(c *fiber.Ctx) error {
	// things needed for new game -> gamename, boardID, playerone, playertwo, Iswinner, CurrentPlayer
	db := database.DBC
	newGame := new(models.NewGameInfo)
	emptyBoard := new(models.GameBoard)
	var firstPlayerID int
	var secondPlayerID int
	var gameBoardID int
	var CreateGame models.Game

	// parse body and check for errors
	if err := c.BodyParser(newGame); err != nil {
		return helperfunctions.ReturnJSONError(c, 500, "Could not parse new game body")
	}

	helperfunctions.FillBoard(emptyBoard, newGame.GameName) // fill board before creating it
	db.Create(&emptyBoard)                                  // create new empty board

	// query for created game board ID
	db.Raw("SELECT id FROM game_boards WHERE board_name = '" + newGame.GameName + "' AND deleted_at IS NULL").Scan(&gameBoardID)
	if gameBoardID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed query for game ID")
	}

	// query for first player ID
	db.Raw("SELECT id FROM tic_users WHERE username = '" + newGame.PlOne + "' AND deleted_at IS NULL").Scan(&firstPlayerID)
	if firstPlayerID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed query for player one id")
	}

	// query for second player ID
	db.Raw("SELECT id FROM tic_users WHERE username = '" + newGame.PlTwo + "' AND deleted_at IS NULL").Scan(&secondPlayerID)
	if secondPlayerID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed query for player two id")
	}

	// form completed game object
	helperfunctions.FillGame(&CreateGame, newGame.GameName, gameBoardID, firstPlayerID, secondPlayerID, newGame.PlOne)

	// check if game object was populated via numeric values
	if CreateGame.BoardID == 0 && CreateGame.PlOneID == 0 {
		helperfunctions.ReturnJSONError(c, 500, "CreateGame Object not properly filled")
	}

	// add game to database
	db.Create(&CreateGame)

	finalMessage := "New Game: '" + newGame.GameName + "' has been created" // message sent back to users
	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}

// this function deletes a game from the db and the corresponding game board
func EndGame(c *fiber.Ctx) error {
	db := database.DBC
	InputGame := new(models.DeleteGame)
	var Game models.Game
	var Board models.GameBoard

	// parse body and check for errors
	if err := c.BodyParser(InputGame); err != nil {
		return helperfunctions.ReturnJSONError(c, 500, "Could not parse delete game body")
	}

	// query for the game ID based on the gamename
	db.Raw("SELECT * FROM games WHERE game_name = '" + InputGame.GameName + "' AND deleted_at IS NULL").Scan(&Game)

	// check if query is valid
	if Game.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed query for delete game")
	}

	// query for the board ID based on the foreign key
	db.Raw("SELECT * FROM game_boards WHERE board_name = '" + InputGame.GameName + "' AND deleted_at IS NULL").Scan(&Board)

	// check if query is valid
	if Board.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed query for delete game board")
	}

	db.Delete(&models.Game{}, Game.ID)       // delete the game
	db.Delete(&models.GameBoard{}, Board.ID) // delete the corresponding game board

	finalMessage := "The Game: '" + InputGame.GameName + "' has been deleted"
	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}

// this function takes a game turn when given a [boardname, username, and a spot(1-9)]
func TakeTurn(c *fiber.Ctx) error {
	// declare variables
	db := database.DBC
	TurnInput := new(models.TurnInput)
	var GameBoard models.GameBoard
	var GameUser models.TicUser
	var TicGame models.Game
	var finalMessage string

	// take in boardname, username and spot(parse body)
	if err := c.BodyParser(TurnInput); err != nil {
		return helperfunctions.ReturnJSONError(c, 500, "Could not parse turn input body")
	}

	// get game info and check if query valid
	db.Raw("SELECT * FROM games WHERE game_name = '" + TurnInput.GameName + "' AND deleted_at IS NULL").Scan(&TicGame)
	if TicGame.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed query for Game information in taketurn")
	}

	// check if game is won
	if TicGame.IsWinner == true {
		return helperfunctions.ReturnJSONError(c, 500, "This game already has a winner")
	}

	// get user info and check if query valid
	db.Raw("SELECT * FROM tic_users WHERE username = '" + TurnInput.Username + "' AND deleted_at IS NULL").Scan(&GameUser)
	if GameUser.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed query for user information in taketurn")
	}

	// check if user in TurnInput is the current player
	if TurnInput.Username != TicGame.CurrentPlayer {
		wrongUserMessage := "It is not your turn: " + TurnInput.Username + "."
		return helperfunctions.ReturnJSONError(c, 500, wrongUserMessage)
	}

	// convert spot input into a fieldname and store in variable?
	correctedField := helperfunctions.ConvertUserInput(TurnInput.TileNum)
	if correctedField == "no_num" {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed user tile num input")
	}
	// query for boardstate and check if spot taken
	db.Raw("SELECT * FROM game_boards WHERE board_name = '" + TurnInput.GameName + "' AND deleted_at IS NULL").Scan(&GameBoard)
	if GameBoard.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed query for game board boardstate")
	}
	checkedSquare := helperfunctions.CheckSquare(correctedField, &GameBoard)
	if checkedSquare == false {
		return helperfunctions.ReturnJSONError(c, 500, "The Spot: '"+correctedField+"' on the board "+GameBoard.BoardName+" is already taken")
	}

	// replace spot with user symbol and check if winner found
	helperfunctions.ReplaceSquare(GameUser.Symbol, correctedField, &GameBoard)
	db.Save(&GameBoard) // save board state before check
	winnerFound := helperfunctions.CheckForWinner(&GameBoard)
	// switch current player if winner not found(PLAYER NOT BEING SWITCHED?)
	if winnerFound == false {
		// function to switch current player
		helperfunctions.SwitchPlayer(int(GameUser.ID), &TicGame)
		finalMessage = "User: " + TurnInput.Username + " has successfully taken their turn."
		db.Save(&TicGame)
	} else {
		// get game object and set is_winner to true and winning_player to current_player
		TicGame.IsWinner = true
		TicGame.WinningPlayer = TicGame.CurrentPlayer
		finalMessage = "User: '" + TurnInput.Username + "' has won the game."
		db.Save(&TicGame)
	}

	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}
