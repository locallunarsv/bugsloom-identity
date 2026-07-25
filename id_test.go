package identity_test

import (
	"testing"

	"github.com/locallunarsv/bugsloom-identity"
)

func TestNewID(t *testing.T) {
	id := identity.New()

	if !id.Valid() {
		t.Fatal("generated ID should be valid")
	}

	if id.String() == "" {
		t.Fatal("ID string should not be empty")
	}
}

func TestParseID(t *testing.T) {
	original := identity.New()

	parsed, err := identity.Parse(original.String())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if original.String() != parsed.String() {
		t.Fatal("parsed ID does not match original")
	}
}

func TestParseInvalidID(t *testing.T) {
	_, err := identity.Parse("invalid-id")

	if err == nil {
		t.Fatal("expected error")
	}
}
