package identity_test

import (
	"encoding/json"
	"testing"

	identity "github.com/locallunarsv/bugsloom-identity"
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

func TestJSONMarshalID(t *testing.T) {
	id := identity.New()

	data, err := json.Marshal(id)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(data) != `"`+id.String()+`"` {
		t.Fatalf(
			"unexpected json output: %s",
			data,
		)
	}
}

func TestJSONUnmarshalID(t *testing.T) {
	original := identity.New()

	data, err := json.Marshal(original)

	if err != nil {
		t.Fatal(err)
	}

	var parsed identity.ID

	err = json.Unmarshal(data, &parsed)

	if err != nil {
		t.Fatal(err)
	}

	if original.String() != parsed.String() {
		t.Fatal("IDs do not match")
	}
}

func TestDatabaseValue(t *testing.T) {
	id := identity.New()

	value, err := id.Value()

	if err != nil {
		t.Fatal(err)
	}

	if value != id.String() {
		t.Fatal("unexpected database value")
	}
}

func TestDatabaseScan(t *testing.T) {
	original := identity.New()

	var scanned identity.ID

	err := scanned.Scan(original.String())

	if err != nil {
		t.Fatal(err)
	}

	if scanned.String() != original.String() {
		t.Fatal("scanned ID mismatch")
	}
}
