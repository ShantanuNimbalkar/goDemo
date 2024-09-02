package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(3, 4)
	expected := 7
	if result != expected {
		t.Errorf("Add(3, 4) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Subtract(10, 4) = %d; want %d", result, expected)
	}
}

func TestIsPositive(t *testing.T) {
	if !IsPositive(5) {
		t.Errorf("IsPositive(5) = false; want true")
	}

	if IsPositive(-5) {
		t.Errorf("IsPositive(-5) = true; want false")
	}

	if IsPositive(0) {
		t.Errorf("IsPositive(0) = true; want false")
	}
}
