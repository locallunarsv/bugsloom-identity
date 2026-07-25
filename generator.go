package identity

import "github.com/google/uuid"

func New() ID {
	return ID{
		value: uuid.Must(uuid.NewV7()),
	}
}
