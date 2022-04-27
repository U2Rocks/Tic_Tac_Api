package routes

import (
	"TicTac_Api/helperfunctions"
	"TicTac_Api/models"

	"TicTac_Api/database"

	"github.com/gofiber/fiber/v2"
)

// this function adds a new user to the game service
func AddUser(c *fiber.Ctx) error {
	db := database.DBC
	userObject := new(models.TicUser)

	if err := c.BodyParser(userObject); err != nil {
		return helperfunctions.ReturnJSONError(c, 500, "Could not parse new user body")
	}

	// manual checks to see if users created are valid
	if userObject.Symbol == "" {
		return helperfunctions.ReturnJSONError(c, 500, "Incoming entry missing symbol")
	}
	if userObject.Username == "" {
		return helperfunctions.ReturnJSONError(c, 500, "Incoming entry missing username")
	}
	if userObject.Password == "" {
		return helperfunctions.ReturnJSONError(c, 500, "Incoming entry missing password")
	}

	db.Create(&userObject)
	finalMessage := "Successfully created the user: " + userObject.Username

	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}

// this function changes a users password when given a username
func ChangePassword(c *fiber.Ctx) error {
	db := database.DBC
	userInput := new(models.ChangeUserInfo)
	var userObject models.TicUser

	if err := c.BodyParser(userInput); err != nil {
		return helperfunctions.ReturnJSONError(c, 500, "Could not parse change password body")
	}

	// get matching query and scan data into struct
	db.Raw("SELECT * FROM tic_users WHERE username = '" + userInput.Username + "' AND password = '" + userInput.Password + "' AND deleted_at IS NULL").Scan(&userObject)

	// check if query worked
	if userObject.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed Query")
	}

	userObject.Password = userInput.NewPassword
	db.Save(&userObject)

	finalMessage := "Password for user: '" + userObject.Username + "' has been changed"

	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}

// this function gets the json for a user object when given a username
func GetUser(c *fiber.Ctx) error {
	db := database.DBC
	userInput := new(models.GetTicUser) // model used to get user search data
	var userObject models.TicUser       // model used to store data from db

	if err := c.BodyParser(userInput); err != nil {
		return helperfunctions.ReturnJSONError(c, 500, "Could not parse get user body")
	}

	// get matching query and scan data into struct
	db.Raw("SELECT * FROM tic_users WHERE username = '" + userInput.Username + "' AND password = '" + userInput.Password + "' AND deleted_at IS NULL").Scan(&userObject)

	// check if query worked
	if userObject.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed Query")
	}

	return c.Status(200).JSON(userObject)
}

// this function gets a slice of all users
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DBC
	var userList []models.TicUser

	db.Find(&userList)
	return c.Status(200).JSON(userList)
}

// this function deletes a user
func DeleteUser(c *fiber.Ctx) error {
	db := database.DBC
	userInput := new(models.GetTicUser)
	var userObject models.TicUser

	if err := c.BodyParser(userInput); err != nil {
		helperfunctions.ReturnJSONError(c, 500, "Could not parse delete user body")
	}

	// get matching query and scan data into struct
	db.Raw("SELECT * FROM tic_users WHERE username = '" + userInput.Username + "' AND password = '" + userInput.Password + "' AND deleted_at IS NULL").Scan(&userObject)

	// check if query worked
	if userObject.ID == 0 {
		return helperfunctions.ReturnJSONError(c, 500, "Malformed Query")
	}

	db.Delete(&models.TicUser{}, userObject.ID) // delete user from database
	finalMessage := "User: " + userObject.Username + " has been successfully deleted"

	return helperfunctions.ReturnJSONResponse(c, 200, finalMessage)
}
