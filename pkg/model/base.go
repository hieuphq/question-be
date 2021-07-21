package model

import (
	"database/sql"
	"time"
)

// Base model
type Base struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *sql.NullTime
}
