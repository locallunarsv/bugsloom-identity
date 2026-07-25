package identity_test

import (
	"errors"
	"testing"

	identity "github.com/locallunarsv/bugsloom-identity"
)

func TestDatabaseValue(t *testing.T) {
	id := identity.New()

	value, err := id.Value()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value != id.String() {
		t.Fatalf(
			"expected %s, got %v",
			id.String(),
			value,
		)
	}
}

func TestDatabaseScanString(t *testing.T) {
	original := identity.New()

	var scanned identity.ID

	err := scanned.Scan(original.String())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if scanned.String() != original.String() {
		t.Fatal("scanned ID does not match")
	}
}

func TestDatabaseScanBytes(t *testing.T) {
	original := identity.New()

	var scanned identity.ID

	err := scanned.Scan([]byte(original.String()))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if scanned.String() != original.String() {
		t.Fatal("scanned ID does not match")
	}
}

func TestDatabaseScanNil(t *testing.T) {
	var id identity.ID

	err := id.Scan(nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !id.IsZero() {
		t.Fatal("expected zero ID")
	}
}

func TestDatabaseScanUnsupportedType(t *testing.T) {
	var id identity.ID

	err := id.Scan(12345)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDatabaseScanInvalidType(t *testing.T) {
	var id identity.ID

	err := id.Scan(12345)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDatabaseValueZeroID(t *testing.T) {
	var id identity.ID

	value, err := id.Value()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value != nil {
		t.Fatalf(
			"expected nil value for zero ID, got %v",
			value,
		)
	}
}

func TestDatabaseScanInvalidString(t *testing.T) {
	var id identity.ID

	err := id.Scan("invalid-id")

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, identity.ErrInvalidID) {
		t.Fatalf(
			"expected ErrInvalidID, got %v",
			err,
		)
	}
}

func TestDatabaseScanInvalidBytes(t *testing.T) {
	var id identity.ID

	err := id.Scan([]byte("invalid-id"))

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, identity.ErrInvalidID) {
		t.Fatalf(
			"expected ErrInvalidID, got %v",
			err,
		)
	}
}