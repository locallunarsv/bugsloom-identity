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
