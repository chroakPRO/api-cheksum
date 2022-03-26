package models

import (
	"github.com/google/uuid"
	_ "os"
)

type PersonalStruct struct {
	UUID      uuid.UUID `db:"uid" json:"uid" validate:"required"`
	Name      string    `db:"name" json:"name" validate:"required"`
	Age       int       `db:"age" json:"age" validate:"required"`
	Portfolio string    `db:"portfolio" json:"portfolio" validate:""`
	Website   string    `db:"website" json:"website" validate:""`
	Employed  bool      `db:"employed" json:"employed" validate:"required"`
	Email     string    `db:"email" json:"email" validate:"required"`
}
