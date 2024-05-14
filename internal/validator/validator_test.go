package validator

import (
	"regexp"
	"testing"
)

func TestNewValidator(t *testing.T) {
	v := New()

	if v == nil {
		t.Errorf("New() returned nil")
	}

	if v.Errors == nil {
		t.Errorf("New() did not initialize the Errors map")
	}

	if len(v.Errors) != 0 {
		t.Errorf("New() did not initialize an empty Errors map")
	}
}

func TestValidator_Valid(t *testing.T) {
	v := New()

	if v.Valid() == false {
		t.Errorf("Expected Valid() to return true for a new Validator instance")
	}

	v.AddError("key", "message")
	if v.Valid() == true {
		t.Errorf("Expected Valid() to return false after adding an error")
	}
}

func TestValidator_AddError(t *testing.T) {
	v := New()
	v.AddError("key", "message")

	if len(v.Errors) != 1 {
		t.Errorf("Expected Errors map to have one entry after AddError, got: %d", len(v.Errors))
	}

	if msg, ok := v.Errors["key"]; !ok || msg != "message" {
		t.Errorf("Expected 'key' with 'message' in the Errors map")
	}
}

func TestPermittedValue(t *testing.T) {
	ok := PermittedValue(2, 1, 2, 3)

	if ok != true {
		t.Errorf("Expected PermittedValue to return true for a permitted value")
	}

	ok = PermittedValue(5, 1, 2, 3)

	if ok != false {
		t.Errorf("Expected PermittedValue to return false for a non-permitted value")
	}
}

func TestMatches(t *testing.T) {
	rx := regexp.MustCompile("^abc.*$")
	ok := Matches("abcdef", rx)

	if ok != true {
		t.Errorf("Expected Matches to return true for matching value")
	}

	ok = Matches("xyz", rx)

	if ok != false {
		t.Errorf("Expected Matches to return false for non-matching value")
	}
}

func TestUnique(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	ok := Unique(slice)

	if ok != true {
		t.Errorf("Expected Unique to return true for slice with unique values")
	}

	slice = []int{1, 2, 3, 4}
	ok = Unique(slice)

	if ok != false {
		t.Errorf("Expected Unique to return false for slice with duplicate values")
	}
}
