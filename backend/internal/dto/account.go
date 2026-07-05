package dto

import (
	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	Name string
}

type UpdateAccountRequest struct {
	Name string
}

type AccountResponse struct {
	ID   uuid.UUID
	Name string
}
