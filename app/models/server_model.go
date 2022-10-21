package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Server struct to describe server object.
type Server struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	UserID       uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Title        string    `db:"title" json:"title" validate:"required,lte=255"`
	Author       string    `db:"author" json:"author" validate:"required,lte=255"`
	ServerStatus int       `db:"server_status" json:"server_status" validate:"required,len=1"`
	ServerAttrs    ServerAttrs `db:"server_attrs" json:"server_attrs" validate:"required,dive"`
}

// ServerAttrs struct to describe server attributes.
type ServerAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}

// Value make the ServerAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (b ServerAttrs) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// Scan make the ServerAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (b *ServerAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &b)
}
