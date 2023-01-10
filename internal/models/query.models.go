package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserProfile struct {
		UserID      uuid.UUID `json:"-"`
		FullName    string    `json:"fullName"`
		Email       string    `json:"email"`
		PhoneNumber string    `json:"phoneNumber"`
		CreatedAt   time.Time `json:"createdAt"`
		CreatedBy   string    `json:"createdBy"`
		UpdatedAt   time.Time `json:"updatedAt"`
		UpdatedBy   string    `json:"updatedBy"`
	}

	AccountParam struct {
		ID        uuid.UUID
		UserName  string    `json:"username"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"createdAt"`
		CreatedBy string    `json:"createdBy"`
		UpdatedAt time.Time `json:"updatedAt"`
		UpdatedBy string    `json:"updatedBy"`
	}
)
