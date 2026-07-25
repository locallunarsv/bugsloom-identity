package identity

import (
	"fmt"

	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

func (id ID) String() string {
	return id.value.String()
}

func (id ID) IsZero() bool {
	return id.value == uuid.Nil
}

func (id ID) Valid() bool {
	return !id.IsZero()
}

func (id ID) GoString() string {
	return fmt.Sprintf("identity.ID(%s)", id.String())
}
