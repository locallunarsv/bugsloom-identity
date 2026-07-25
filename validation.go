package identity

import "github.com/google/uuid"

func validateUUID(id uuid.UUID) error {
	if id.Version() != 7 {
		return ErrInvalidID
	}

	return nil
}