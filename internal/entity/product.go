package entity

import (
	"errors"
	"time"

	"github.com/jeffersonbraster/apigo/pkg/entity"
)

var (
	ErrIdIsRequired = errors.New("id is required")
	ErrInvalidId = errors.New("invalid id")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice = errors.New("invalid price")
)

type Product struct {
	ID          entity.ID 	 `json:"id"`
	Name        string  `json:"name"`
	Price int `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	p := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdIsRequired
	}
	if _, err := entity.PaserID(p.ID.String()); err != nil {
		return ErrInvalidId
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
