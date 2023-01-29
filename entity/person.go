package entity

import (
	"github.com/google/uuid"
)

// Person is an entity that represents a person in all domains
type Person struct {
	// ID is the identifier of the Entity, The ID is shared for all sub-domains
	ID   uuid.UUID
	Name string
	Age  int
}
