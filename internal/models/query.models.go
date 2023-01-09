package models

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type (
	UserProfile struct {
		ID          uuid.UUID `json:"-"`
		UserID      uuid.UUID `json:"-"`
		FullName    string    `json:"fullName"`
		Email       string    `json:"email"`
		PhoneNumber string    `json:"phoneNumber"`
		CreatedAt   time.Time `json:"createdAt"`
		CreatedBy   string    `json:"createdBy"`
		UpdatedAt   time.Time `json:"updatedAt"`
		UpdatedBy   string    `json:"updatedBy"`
	}
)
