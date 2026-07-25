package identity_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	identity "github.com/locallunarsv/bugsloom-identity"
)

func TestParseID(t *testing.T) {
	original := identity.New()

	parsed, err := identity.Parse(original.String())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if parsed.String() != original.String() {
		t.Fatal("parsed ID does not match original")
	}
}

func TestParseEmptyID(t *testing.T) {
	_, err := identity.Parse("")

	if !errors.Is(err, identity.ErrEmptyID) {
		t.Fatal("expected ErrEmptyID")
	}
}

func TestParseInvalidID(t *testing.T) {
	_, err := identity.Parse("invalid-id")

	if !errors.Is(err, identity.ErrInvalidID) {
		t.Fatal("expected ErrInvalidID")
	}
}

func TestParseRejectNonV7UUID(t *testing.T) {
	v4 := uuid.New()

	_, err := identity.Parse(v4.String())

	if !errors.Is(err, identity.ErrInvalidID) {
		t.Fatal("expected invalid ID error")
	}
}