package api

import (
	"paste/routes"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ReadOneController is a controller that can be accessed through /api/post/:ID
func ReadOneController(c *fiber.Ctx) error {
	connection := CreateConnection()
	ID, err := strconv.ParseInt(c.Params("ID"), 10, 32)
	defer connection.Connection.Close()

	if err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Internal Server Error - Invalid Parameter",
			Success: false,
			Data:    nil,
		})
	}

	post, err := connection.getPost(ID)

	if err != nil {
		return c.JSON(routes.HTTPResponse{
			Message: "Internal Server Error",
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(routes.HTTPResponse{
		Message: nil,
		Success: true,
		Data:    post,
	})
}
