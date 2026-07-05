package dto

import "github.com/google/uuid"

type CategoryResponse struct {
	ID   uuid.UUID
	Name string
}
