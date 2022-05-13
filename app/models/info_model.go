package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Info struct to describe Info object.
type Info struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	UserID     uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Name       string    `db:"name" json:"name" validate:"required,lte=255"`
	Portfolio  string    `db:"portfolio" json:"portfolio" validate:"required,lte=255"`
	InfoStatus int       `db:"Info_status" json:"Info_status" validate:"required,len=1"`
	InfoAttrs  InfoAttrs `db:"Info_attrs" json:"Info_attrs" validate:"required,dive"`
}

// InfoAttrs struct to describe Info attributes.
type InfoAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

// Value make the InfoAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (b InfoAttrs) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// Scan make the InfoAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (b *InfoAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &b)
}
