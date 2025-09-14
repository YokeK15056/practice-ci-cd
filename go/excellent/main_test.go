package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(3)
	if result != "Odd" {
		t.Errorf("Expected 'Odd', but got '%s'", result)
	}
}