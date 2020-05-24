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

// Import loads the filename content, json into the store. Returning
// number of items currently stored.
func (s *Store) Import(filename string) (int, error) {
	return 0, fmt.Errorf("TODO")
}

type Product struct {
	Name     string
	Category string
}
