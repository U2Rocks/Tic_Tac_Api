package helperfunctions

import (
	"TicTac_Api/models"

	"github.com/gofiber/fiber/v2"
)

// this function returns a customized error and needs [context, errorcode, message]
func ReturnJSONError(c *fiber.Ctx, errorcode uint, message string) error {
	Response := models.ResponseObject{StatusCode: errorcode, Message: message}
	intcode := int(errorcode)
	return c.Status(intcode).JSON(Response)
}

// wrapper function that checks if error is nil and then executes return JSON error
func CheckError(c *fiber.Ctx, errorcode uint, err error, message string) {
	if err != nil {
		ReturnJSONError(c, errorcode, message)
	}
}

// this function returns a JSON response to the user needs [context, statuscode, message]
func ReturnJSONResponse(c *fiber.Ctx, statuscode uint, message string) error {
	PositiveResponse := models.ResponseObject{StatusCode: statuscode, Message: message}
	intstatus := int(statuscode)
	return c.Status(intstatus).JSON(PositiveResponse)
}
