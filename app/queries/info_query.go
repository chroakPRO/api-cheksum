package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/models"
)

// InfoQueries struct for queries from Info model.
type InfoQueries struct {
	*sqlx.DB
}

// GetInfo method for getting all Info.
func (q *InfoQueries) GetAllInfo() ([]models.Info, error) {
	// Define Info variable.
	Info := []models.Info{}

	// Define query string.
	query := `SELECT * FROM Info`

	// Send query to database.
	err := q.Get(&Info, query)
	if err != nil {
		// Return empty object and error.
		return Info, err
	}

	// Return query result.
	return Info, nil
}

// GetInfo method for getting one Info by given ID.
func (q *InfoQueries) GetInfo(id uuid.UUID) (models.Info, error) {
	// Define Info variable.
	Info := models.Info{}

	// Define query string.
	query := `SELECT * FROM Info WHERE id = $1`

	// Send query to database.
	err := q.Get(&Info, query, id)
	if err != nil {
		// Return empty object and error.
		return Info, err
	}

	// Return query result.
	return Info, nil
}

// CreateInfo method for creating Info by given Info object.
func (q *InfoQueries) CreateInfo(b *models.Info) error {
	// Define query string.
	query := `INSERT INTO Info VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Name, b.Portfolio, b.InfoStatus, b.InfoAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateInfo method for updating Info by given Info object.
func (q *InfoQueries) UpdateInfo(id uuid.UUID, b *models.Info) error {
	// Define query string.
	query := `UPDATE Info SET updated_at = $2, title = $3, author = $4, Info_status = $5, Info_attrs = $6 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id, b.UpdatedAt, b.UserID, b.Name, b.InfoStatus, b.InfoAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteInfo method for delete Info by given ID.
func (q *InfoQueries) DeleteInfo(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM Info WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
