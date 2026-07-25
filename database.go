package identity

import (
	"database/sql/driver"
	"fmt"
)

func (id ID) Value() (driver.Value, error) {
	if !id.Valid() {
		return nil, nil
	}

	return id.String(), nil
}

func (id *ID) Scan(value any) error {
	switch v := value.(type) {

	case nil:
		*id = ID{}
		return nil

	case string:
		parsed, err := Parse(v)

		if err != nil {
			return err
		}

		*id = parsed
		return nil

	case []byte:
		parsed, err := Parse(string(v))

		if err != nil {
			return err
		}

		*id = parsed
		return nil

	default:
		return fmt.Errorf(
			"identity: cannot scan type %T",
			value,
		)
	}
}