package model

import (
	"time"

	"github.com/influxdata/influxdb/uuid"
)

// User model
type User struct {
	ID               uint       `json:"id"`
	UUID             uuid.UUID  `json:"uuid"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	MasterPassword   string     `json:"master_password"`
	Secret           string     `json:"secret"`
	Schema           string     `json:"schema"`
	Role             string     `json:"role"`
	ConfirmationCode string     `json:"confirmation_code"`
	EmailVerifiedAt  time.Time  `json:"email_verified_at"`
}
