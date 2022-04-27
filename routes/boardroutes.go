package routes

import (
	"TicTac_Api/helperfunctions"

	"TicTac_Api/models"

	"TicTac_Api/database"

	"github.com/gofiber/fiber/v2"
)

// this route returns the current state of the board and takes in [board_name]
func BoardState(c *fiber.Ctx) error {
	db := database.DBC
	Status := new(models.GameStatus)
	var GameBoard models.GameBoard

	// parse body and check for errors
	if err := c.BodyParser(Status); err != nil {
		return helperfunctions.ReturnJSONError(c, 500, "Could not parse board state body")
	}

	// query info on game board
	db.Raw("SELECT * FROM game_boards WHERE board_name = '" + Status.GameName + "' AND deleted_at IS NULL").Scan(&GameBoard)
	// check if query is valid
	if GameBoard.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed query for gameboard status")
	}

	return c.Status(200).JSON(GameBoard)
}
