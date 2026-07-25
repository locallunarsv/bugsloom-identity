package identity

import "github.com/google/uuid"

func Parse(value string) (ID, error) {
	if value == "" {
		return ID{}, ErrEmptyID
	}

	id, err := uuid.Parse(value)

	if err != nil {
		return ID{}, ErrInvalidID
	}

	if err := validateUUID(id); err != nil {
		return ID{}, err
	}

	return ID{
		value: id,
	}, nil
}