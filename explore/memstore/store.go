package memstore

import "fmt"

func NewStore() *Store {
	return &Store{
		products: make([]*Product, 0),
	}
}

type Store struct {
	products []*Product
}

func (s *Store) Import(filename string) (int, error) {
	return 0, fmt.Errorf("TODO")
}

type Product struct {
	Name     string
	Category string
}
