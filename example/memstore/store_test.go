package memstore

/*
  Import products into a memory collection
*/

import (
	"testing"
)

func xTest_import(t *testing.T) {
	s := NewStore()
	n, err := s.Import("./products.json")
	if err != nil {
		t.Fatal(err, "tip! go run ../downloads > products.json")
	}
	if n == 0 {
		t.Fatal("no products imported")
	}
	if s.products[0].Name == "" {
		t.Error("Name is empty")
	}
	if s.products[0].Category == "" {
		t.Error("Category is empty")
	}
}
