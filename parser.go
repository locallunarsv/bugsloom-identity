package identity

import "github.com/google/uuid"

func Parse(value string) (ID, error) {
	id, err := uuid.Parse(value)

	if err != nil {
		return ID{}, err
	}

	return ID{
		value: id,
	}, nil
}
