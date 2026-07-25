package identity_test

import (
	"testing"

	identity "github.com/locallunarsv/bugsloom-identity"
)

func BenchmarkNewID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		identity.New()
	}
}