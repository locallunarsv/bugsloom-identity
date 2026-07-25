package identity_test

import (
	"fmt"
	"testing"

	identity "github.com/locallunarsv/bugsloom-identity"
)

func TestNewID(t *testing.T) {
	id := identity.New()

	if !id.Valid() {
		t.Fatal("new ID should be valid")
	}

	if id.IsZero() {
		t.Fatal("new ID should not be zero")
	}

	if id.String() == "" {
		t.Fatal("ID string should not be empty")
	}
}

func TestZeroID(t *testing.T) {
	var id identity.ID

	if id.Valid() {
		t.Fatal("zero ID should be invalid")
	}

	if !id.IsZero() {
		t.Fatal("zero ID should be zero")
	}
}

func TestGoString(t *testing.T) {
	id := identity.New()

	got := fmt.Sprintf("%#v", id)

	expected := "identity.ID(" + id.String() + ")"

	if got != expected {
		t.Fatalf(
			"expected %s, got %s",
			expected,
			got,
		)
	}
}