package identity

import "encoding/json"

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *ID) UnmarshalJSON(data []byte) error {
	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	parsed, err := Parse(value)

	if err != nil {
		return err
	}

	*id = parsed

	return nil
}