package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		value := add(2, 2)

		if value != 4 {
			t.Errorf("expected: %d, got: %d", 4, value)
		}
	})
}