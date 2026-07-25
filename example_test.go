package identity_test

import (
	"fmt"

	identity "github.com/locallunarsv/bugsloom-identity"
)

func ExampleNew() {
	id := identity.New()

	fmt.Println(id.Valid())

	// Output:
	// true
}