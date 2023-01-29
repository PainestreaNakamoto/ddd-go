package aggregate

import (
	"errors"

	"github.com/PainestreaNakamoto/ddd-go/entity"
	"github.com/PainestreaNakamoto/ddd-go/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

type Customer struct {
	person      *entity.Person
	product     []*entity.Item
	transaction []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}
	return Customer{
		person:      person,
		product:     make([]*entity.Item, 0),
		transaction: make([]valueobject.Transaction, 0),
	}, nil
}
