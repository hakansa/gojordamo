package model

import (
	"time"

	"github.com/influxdata/influxdb/uuid"
)

// User model
type User struct {
	ID               uint      `json:"id"`
	UUID             uuid.UUID `json:"uuid"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Secret           string    `json:"secret"`
	Schema           string    `json:"schema"`
	Role             string    `json:"role"`
	ConfirmationCode string    `json:"confirmation_code"`
	EmailVerifiedAt  time.Time `json:"email_verified_at"`
}

// GetUsersResults collects the results of the GetUsers call: the list of Users matching
// the HeaderFilterOptions, and the TotalCount of the matching users before paging was applied.
type GetUsersResults struct {
	TotalCount int    `json:"total_count"`
	PageCount  int    `json:"page_count"`
	HasMore    bool   `json:"has_more"`
	Items      []User `json:"items"`
}

// Store defines the methods the ServiceImpl needs from the interfaceStore.
type Store interface {

	// GetUsers returns filtered users and the total count before paging.
	GetUsers(options FilterOptions) (*GetUsersResults, error)

	// CreateUser creates a new user
	CreateUser(user *User) (*User, error)

	// UpdateUser updates the given user
	UpdateUser(user *User) (*User, error)

	// GetUser gets a user by ID.
	GetUser(id int) (*User, error)

	// NukeDB removes all incident related data.
	NukeDB() error
}
