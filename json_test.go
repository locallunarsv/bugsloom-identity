package identity_test

import (
	"encoding/json"
	"testing"

	identity "github.com/locallunarsv/bugsloom-identity"
)

func TestMarshalJSON(t *testing.T) {
	id := identity.New()

	data, err := json.Marshal(id)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `"` + id.String() + `"`

	if string(data) != expected {
		t.Fatalf(
			"expected %s, got %s",
			expected,
			data,
		)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	original := identity.New()

	data, err := json.Marshal(original)

	if err != nil {
		t.Fatal(err)
	}

	var parsed identity.ID

	err = json.Unmarshal(data, &parsed)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if parsed.String() != original.String() {
		t.Fatal("unmarshaled ID does not match")
	}
}

func TestUnmarshalInvalidJSON(t *testing.T) {
	var id identity.ID

	err := json.Unmarshal(
		[]byte(`"invalid-id"`),
		&id,
	)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestUnmarshalInvalidJSONType(t *testing.T) {
	var id identity.ID

	err := json.Unmarshal(
		[]byte(`123`),
		&id,
	)

	if err == nil {
		t.Fatal("expected error")
	}
}