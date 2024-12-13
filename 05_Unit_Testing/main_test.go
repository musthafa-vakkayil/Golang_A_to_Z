package main

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	width, length := 3, 5
	expected := 15

	area, err := CalculateArea(width, length)

	if area != expected {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}

func TestCalculateAreaNegativeWidth(t *testing.T) {
	width, length := -4, 6

	area, err := CalculateArea(width, length)

	if area != 0 {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}
}

func TestCalculateAreaNegativeLength(t *testing.T) {
	width, length := 5, -7

	area, err := CalculateArea(width, length)

	if area != 0 {
		t.Errorf("wrong value for 'area', expected: %v, got %v", 0, area)
	}
	if err == nil {
		t.Error("expected an error")
	}
}

func TestCalculator(t *testing.T) {
	t.Parallel()
	t.SkipNow()
	var c Calculator
	t.Run("Add", func(t *testing.T) {
		expected := 5.0

		result := c.Add(2, 3)

		if result != expected {
			t.Fail()
		}
	})

	t.Run("Subtract", func(t *testing.T) {
		expected := 2.0

		result := c.Subtract(5, 3)

		if result != expected {
			t.Fail()
		}
	})

	t.Run("Multiply", func(t *testing.T) {
		expected := 6.0

		result := c.Multiply(2, 3)

		if result != expected {
			t.Fail()
		}
	})

	t.Run("Divide", func(t *testing.T) {

		t.Run("Valid", func(t *testing.T) {
			expected := 2.0

			result, err := c.Divide(6, 3)

			if err != nil {
				t.Fail()
			}

			if result != expected {
				t.Fail()
			}
		})

		t.Run("ByZero", func(t *testing.T) {
			_, err := c.Divide(6, 0)

			if err == nil {
				t.Fail()
			}
		})
	})
}
