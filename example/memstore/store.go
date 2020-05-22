package memstore

import (
	"encoding/json"
	"os"
)

func NewStore() *Store {
	return &Store{
		products: make([]*Product, 0),
	}
}

type Store struct {
	products []*Product
}

func (s *Store) Import(filename string) (int, error) {
	in, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	dec := json.NewDecoder(in)
	err = dec.Decode(&s.products)
	return len(s.products), err
}

type Product struct {
	Name     string
	Category string
}
