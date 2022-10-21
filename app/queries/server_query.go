package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/models"
)

// ServerQueries struct for queries from Server model.
type ServerQueries struct {
	*sqlx.DB
}

// GetServers method for getting all servers.
func (q *ServerQueries) GetServers() ([]models.Server, error) {
	// Define servers variable.
	servers := []models.Server{}

	// Define query string.
	query := `SELECT * FROM servers`

	// Send query to database.
	err := q.Get(&servers, query)
	if err != nil {
		// Return empty object and error.
		return servers, err
	}

	// Return query result.
	return servers, nil
}

// GetServer method for getting one server by given ID.
func (q *ServerQueries) GetServer(id uuid.UUID) (models.Server, error) {
	// Define server variable.
	server := models.Server{}

	// Define query string.
	query := `SELECT * FROM servers WHERE id = $1`

	// Send query to database.
	err := q.Get(&server, query, id)
	if err != nil {
		// Return empty object and error.
		return server, err
	}

	// Return query result.
	return server, nil
}

// CreateServer method for creating server by given Server object.
func (q *ServerQueries) CreateServer(b *models.Server) error {
	// Define query string.
	query := `INSERT INTO servers VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Title, b.Author, b.ServerStatus, b.ServerAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateServer method for updating server by given Server object.
func (q *ServerQueries) UpdateServer(id uuid.UUID, b *models.Server) error {
	// Define query string.
	query := `UPDATE servers SET updated_at = $2, title = $3, author = $4, server_status = $5, server_attrs = $6 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id, b.UpdatedAt, b.Title, b.Author, b.ServerStatus, b.ServerAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteServer method for delete server by given ID.
func (q *ServerQueries) DeleteServer(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM servers WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
