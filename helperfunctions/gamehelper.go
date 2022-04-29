package helperfunctions

import (
	"TicTac_Api/database"
	"TicTac_Api/models"
	"strconv"
)

// this function populates an empty gameboard
func FillBoard(gb *models.GameBoard, name string) {
	gb.BoardName = name
	gb.SqOne = " "
	gb.SqTwo = " "
	gb.SqThree = " "
	gb.SqFour = " "
	gb.SqFive = " "
	gb.SqSix = " "
	gb.SqSeven = " "
	gb.SqEight = " "
	gb.SqNine = " "
}

// this function populates a new game entry
func FillGame(g *models.Game, gamename string, boardnum int, plOneID int, plTwoID int, firstplayer string) {
	g.GameName = gamename
	g.BoardID = boardnum
	g.PlOneID = plOneID
	g.PlTwoID = plTwoID
	g.IsWinner = false
	g.CurrentPlayer = firstplayer
}

// this function converts a number to a proper field name
func ConvertUserInput(input int) string {
	if input == 1 {
		return "sq_one"
	}
	if input == 2 {
		return "sq_two"
	}
	if input == 3 {
		return "sq_three"
	}
	if input == 4 {
		return "sq_four"
	}
	if input == 5 {
		return "sq_five"
	}
	if input == 6 {
		return "sq_six"
	}
	if input == 7 {
		return "sq_seven"
	}
	if input == 8 {
		return "sq_eight"
	}
	if input == 9 {
		return "sq_nine"
	}
	return "no_num"
}

// this function replaces a square in a board
func ReplaceSquare(symbol string, spot_to_change string, G *models.GameBoard) {
	if spot_to_change == "sq_one" {
		G.SqOne = symbol
	}
	if spot_to_change == "sq_two" {
		G.SqTwo = symbol
	}
	if spot_to_change == "sq_three" {
		G.SqThree = symbol
	}
	if spot_to_change == "sq_four" {
		G.SqFour = symbol
	}
	if spot_to_change == "sq_five" {
		G.SqFive = symbol
	}
	if spot_to_change == "sq_six" {
		G.SqSix = symbol
	}
	if spot_to_change == "sq_seven" {
		G.SqSeven = symbol
	}
	if spot_to_change == "sq_eight" {
		G.SqEight = symbol
	}
	if spot_to_change == "sq_nine" {
		G.SqNine = symbol
	}
}

// this function checks if a spot on the board is taken
func CheckSquare(spot_to_change string, G *models.GameBoard) bool {
	// these conditionals check if the square is a blank space
	if spot_to_change == "sq_one" {
		if G.SqOne != " " {
			return false
		}
		return true
	}

	if spot_to_change == "sq_two" {
		if G.SqTwo != " " {
			return false
		}
		return true
	}

	if spot_to_change == "sq_three" {
		if G.SqThree != " " {
			return false
		}
		return true
	}

	if spot_to_change == "sq_four" {
		if G.SqFour != " " {
			return false
		}
		return true
	}

	if spot_to_change == "sq_five" {
		if G.SqFive != " " {
			return false
		}
		return true
	}

	if spot_to_change == "sq_six" {
		if G.SqSix != " " {
			return false
		}
		return true
	}

	if spot_to_change == "sq_seven" {
		if G.SqSeven != " " {
			return false
		}
		return true
	}

	if spot_to_change == "sq_eight" {
		if G.SqEight != " " {
			return false
		}
		return true
	}

	if spot_to_change == "sq_nine" {
		if G.SqNine != " " {
			return false
		}
		return true
	}

	return false
}

// this function checks if a winner has been found(need to exclude " " from a winning value)
func CheckForWinner(G *models.GameBoard) bool {
	// check for horizonal wins
	if G.SqOne == G.SqTwo && G.SqOne == G.SqThree && G.SqOne != " " {
		return true
	}
	if G.SqFour == G.SqFive && G.SqFour == G.SqSix && G.SqFour != " " {
		return true
	}
	if G.SqSeven == G.SqEight && G.SqSeven == G.SqNine && G.SqSeven != " " {
		return true
	}

	// check for vertical wins
	if G.SqOne == G.SqFour && G.SqOne == G.SqSeven && G.SqOne != " " {
		return true
	}
	if G.SqTwo == G.SqFive && G.SqTwo == G.SqEight && G.SqTwo != " " {
		return true
	}
	if G.SqThree == G.SqSix && G.SqThree == G.SqNine && G.SqThree != " " {
		return true
	}

	// check for diagonal wins
	if G.SqOne == G.SqFive && G.SqOne == G.SqNine && G.SqOne != " " {
		return true
	}
	if G.SqThree == G.SqFive && G.SqThree == G.SqSeven && G.SqThree != " " {
		return true
	}

	return false
}

// this function switches the current player for a game object(if bad entry -> switches to plOne by default)
func SwitchPlayer(inputID int, G *models.Game) {
	var newUsername string
	db := database.DBC
	if inputID == G.PlOneID { // switch current player from plOne -> plTwo
		newID := strconv.FormatInt(int64(G.PlTwoID), 10)
		db.Raw("SELECT username FROM tic_users WHERE id = " + newID + " AND deleted_at IS NULL").Scan(&newUsername)
		G.CurrentPlayer = newUsername
	} else { // switch current player from plTwo -> plOne
		newID := strconv.FormatInt(int64(G.PlOneID), 10)
		db.Raw("SELECT username FROM tic_users WHERE id = " + newID + " AND deleted_at IS NULL").Scan(&newUsername)
		G.CurrentPlayer = newUsername
	}
}

func CheckForTie(G *models.GameBoard) bool {
	// check if all squares are not equal to a space
	if G.SqOne != " " && G.SqTwo != " " && G.SqThree != " " && G.SqFour != " " && G.SqFive != " " && G.SqSix != " " && G.SqSeven != " " && G.SqEight != " " && G.SqNine != " " {
		// return true if that condition met
		return true
	}
	// return true if that condition met
	return false
}
