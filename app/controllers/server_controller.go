package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/models"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/utils"
	"github.com/koddr/tutorial-go-fiber-rest-api/platform/database"
)

// GetServers func gets all exists servers.
// @Description Get all exists servers.
// @Summary get all exists servers
// @Tags Servers
// @Accept json
// @Produce json
// @Success 200 {array} models.Server
// @Router /v1/servers [get]
func GetServers(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all servers.
	servers, err := db.GetServers()
	if err != nil {
		// Return, if servers not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "servers were not found",
			"count":   0,
			"servers": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(servers),
		"servers": servers,
	})
}

// GetServer func gets server by given ID or 404 error.
// @Description Get server by given ID.
// @Summary get server by given ID
// @Tags Server
// @Accept json
// @Produce json
// @Param id path string true "Server ID"
// @Success 200 {object} models.Server
// @Router /v1/server/{id} [get]
func GetServer(c *fiber.Ctx) error {
	// Catch server ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get server by ID.
	server, err := db.GetServer(id)
	if err != nil {
		// Return, if server not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "server with the given ID is not found",
			"server": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"server": server,
	})
}

// CreateServer func for creates a new server.
// @Description Create a new server.
// @Summary create a new server
// @Tags Server
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param server_attrs body models.ServerAttrs true "Server attributes"
// @Success 200 {object} models.Server
// @Security ApiKeyAuth
// @Router /v1/server [post]
func CreateServer(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current server.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Server struct
	server := &models.Server{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(server); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Server model.
	validate := utils.NewValidator()

	// Set initialized default data for server:
	server.ID = uuid.New()
	server.CreatedAt = time.Now()
	server.ServerStatus = 1 // 0 == draft, 1 == active

	// Validate server fields.
	if err := validate.Struct(server); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete server by given ID.
	if err := db.CreateServer(server); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"server": server,
	})
}

// UpdateServer func for updates server by given ID.
// @Description Update server.
// @Summary update server
// @Tags Server
// @Accept json
// @Produce json
// @Param id body string true "Server ID"
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param server_status body integer true "Server status"
// @Param server_attrs body models.ServerAttrs true "Server attributes"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/server [put]
func UpdateServer(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current server.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Server struct
	server := &models.Server{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(server); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if server with given ID is exists.
	foundedServer, err := db.GetServer(server.ID)
	if err != nil {
		// Return status 404 and server not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "server with this ID not found",
		})
	}

	// Set initialized default data for server:
	server.UpdatedAt = time.Now()

	// Create a new validator for a Server model.
	validate := utils.NewValidator()

	// Validate server fields.
	if err := validate.Struct(server); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Update server by given ID.
	if err := db.UpdateServer(foundedServer.ID, server); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201.
	return c.SendStatus(fiber.StatusCreated)
}

// DeleteServer func for deletes server by given ID.
// @Description Delete server by given ID.
// @Summary delete server by given ID
// @Tags Server
// @Accept json
// @Produce json
// @Param id body string true "Server ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/server [delete]
func DeleteServer(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current server.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Server struct
	server := &models.Server{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(server); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Server model.
	validate := utils.NewValidator()

	// Validate only one server field ID.
	if err := validate.StructPartial(server, "id"); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if server with given ID is exists.
	foundedServer, err := db.GetServer(server.ID)
	if err != nil {
		// Return status 404 and server not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "server with this ID not found",
		})
	}

	// Delete server by given ID.
	if err := db.DeleteServer(foundedServer.ID); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
