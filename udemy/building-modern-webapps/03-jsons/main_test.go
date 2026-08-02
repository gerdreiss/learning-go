package main

import (
	"testing"
	"testing/quick"
)

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Expected no error but got one")
			}
		}
		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

func TestDivisionProperty(t *testing.T) {
	property := func(a, b float32) bool {
		if b == 0.0 {
			res, err := divide(a, b)
			return res == 0.0 && err != nil
		}
		res, err := divide(a, b)
		return res == a/b && err == nil
	}

	// quick.Check runs the property with many random inputs
	// The 'nil' config uses defaults (100 iterations by default)
	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestDivideError(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("Got no error when we should have")
	}
}
