package main

import "testing"

func Test(t *testing.T) {
	c := &cli{}
	if err := c.run(); err == nil {
		t.Error("should fail")
	}
}
