package dtos

import (
	"time"

	"github.com/google/uuid"
)

type (
	Account struct {
		ID        uuid.UUID `json:"-"`
		UserName  string    `json:"username"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"createdAt"`
		CreatedBy string    `json:"createdBy"`
		UpdatedAt time.Time `json:"updatedAt"`
		UpdatedBy string    `json:"updatedBy"`
	}

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
