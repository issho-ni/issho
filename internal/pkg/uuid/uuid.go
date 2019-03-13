package uuid

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
)

// UUID implements additional interfaces atop the uuid.UUID type.
type UUID struct {
	uuid.UUID
}

// New generates a new UUID.
func New() UUID {
	return UUID{uuid.New()}
}

// UnmarshalGQL implements the graphql.Unmarshal interface.
func (u *UUID) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("Value for unmarshaling was not a string: %v", v)
	}

	parsed, err := uuid.Parse(str)
	if err != nil {
		return err
	}

	u.UUID = parsed
	return nil
}

// MarshalGQL implements the graphql.Marshal interface.
func (u UUID) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(u.String())))
}

// Size is required to implement the proto.Marshaler interface.
func (u *UUID) Size() int {
	if u == nil || len(u.UUID) == 0 {
		return 0
	}
	return 16
}

// MarshalTo is required to implement the proto.Marshaler interface.
func (u *UUID) MarshalTo(data []byte) (int, error) {
	if u == nil || len(u.UUID) == 0 {
		return 0, nil
	}
	copy(data, u.UUID[:])
	return 16, nil
}

// Unmarshal is required to implement the proto.Marshaler interface.
func (u *UUID) Unmarshal(data []byte) error {
	if len(data) == 0 {
		u = nil
		return nil
	}

	uid, err := uuid.FromBytes(data)
	if err != nil {
		return err
	}

	u.UUID = uid
	return nil
}
