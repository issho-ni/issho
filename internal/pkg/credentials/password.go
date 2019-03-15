package credentials

import (
	"fmt"
	"io"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// Password is a bcrypt-hashed credential value
type Password []byte

// UnmarshalGQL implements the graphql.Unmarshal interface.
func (p *Password) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("Value for unmarshaling was not a string")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	pass := Password(password)
	*p = pass
	return nil
}

// MarshalGQL implements the graphql.Marshal interface.
func (p Password) MarshalGQL(w io.Writer) {}

// Size is required to implement the proto.Marshaler interface.
func (p *Password) Size() int {
	if p == nil {
		return 0
	}
	return len(*p)
}

// MarshalTo is required to implement the proto.Marshaler interface.
func (p Password) MarshalTo(data []byte) (int, error) {
	if p == nil {
		return 0, nil
	}
	copy(data, p)
	return len(p), nil
}

// Unmarshal is required to implement the proto.Marshaler interface.
func (p *Password) Unmarshal(data []byte) error {
	pass := Password(data)
	*p = pass
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (p Password) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(string(p))), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (p *Password) UnmarshalJSON(data []byte) error {
	return p.Unmarshal(data)
}
