package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/models"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/utils"
	"github.com/koddr/tutorial-go-fiber-rest-api/platform/database"
)

// GetInfo func gets Info by given ID or 404 error.
// @Description Get Info by given ID.
// @Summary get Info by given ID
// @Tags Info
// @Accept json
// @Produce json
// @Param id path string true "Info ID"
// @Success 200 {object} models.Info
// @Router /v1/info/{id} [get]
func GetInfo(c *fiber.Ctx) error {
	// Catch Info ID from URL.
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

	// Get Info by ID.
	Info, err := db.GetInfo(id)
	if err != nil {
		// Return, if Info not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Info with the given ID is not found",
			"Info":  nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"Info":  Info,
	})
}

func GetInfoByID(c *fiber.Ctx) error {
	// Catch Info ID from URL.
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

	// Get Info by ID.
	Info, err := db.GetInfo(id)
	if err != nil {
		// Return, if Info not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Info with the given ID is not found",
			"Info":  nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"Info":  Info,
	})
}

// CreateInfo func for creates a new Info.
// @Description Create a new Info.
// @Summary create a new Info
// @Tags Info
// @Accept json
// @Produce json
// @Param Name body string true "name"
// @Param Portfolio body string true "website"
// @Param InfoStatus body string true "currently employed or"
// @Param Info_attrs body models.InfoAttrs true "Info attributes"
// @Success 200 {object} models.Info
// @Security ApiKeyAuth
// @Router /v1/info [post]
func CreateInfo(c *fiber.Ctx) error {
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

	// Set expiration time from JWT data of current Info.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Info struct
	Info := &models.Info{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(Info); err != nil {
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

	// Create a new validator for a Info model.
	validate := utils.NewValidator()

	// Set initialized default data for Info:
	Info.ID = uuid.New()
	Info.CreatedAt = time.Now()
	Info.InfoStatus = 1 // 0 == draft, 1 == active

	// Validate Info fields.
	if err := validate.Struct(Info); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete Info by given ID.
	if err := db.CreateInfo(Info); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"Info":  Info,
	})
}

// UpdateInfo func for updates Info by given ID.
// @Description Update Info.
// @Summary update Info
// @Tags Info
// @Accept json
// @Produce json
// @Param id body string true "Info ID"
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param Info_status body integer true "Info status"
// @Param Info_attrs body models.InfoAttrs true "Info attributes"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/info [put]
func UpdateInfo(c *fiber.Ctx) error {
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

	// Set expiration time from JWT data of current Info.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Info struct
	Info := &models.Info{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(Info); err != nil {
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

	// Checking, if Info with given ID is exists.
	foundedInfo, err := db.GetInfo(Info.ID)
	if err != nil {
		// Return status 404 and Info not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Info with this ID not found",
		})
	}

	// Set initialized default data for Info:
	Info.UpdatedAt = time.Now()

	// Create a new validator for a Info model.
	validate := utils.NewValidator()

	// Validate Info fields.
	if err := validate.Struct(Info); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Update Info by given ID.
	if err := db.UpdateInfo(foundedInfo.ID, Info); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201.
	return c.SendStatus(fiber.StatusCreated)
}

// DeleteInfo func for deletes Info by given ID.
// @Description Delete Info by given ID.
// @Summary delete Info by given ID
// @Tags Info
// @Accept json
// @Produce json
// @Param id body string true "Info ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/info [delete]
func DeleteInfo(c *fiber.Ctx) error {
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

	// Set expiration time from JWT data of current Info.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Info struct
	Info := &models.Info{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(Info); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Info model.
	validate := utils.NewValidator()

	// Validate only one Info field ID.
	if err := validate.StructPartial(Info, "id"); err != nil {
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

	// Checking, if Info with given ID is exists.
	foundedInfo, err := db.GetInfo(Info.ID)
	if err != nil {
		// Return status 404 and Info not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Info with this ID not found",
		})
	}

	// Delete Info by given ID.
	if err := db.DeleteInfo(foundedInfo.ID); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
