package entity

import (
	"github.com/google/uuid"
)

type Item struct {
	// ID is the identifier of the Entity, The ID is shared for all sub-domains
	ID          uuid.UUID
	Name        string
	Description int
}
