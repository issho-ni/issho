package uuid

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Parse decodes s into a UUID or returns an error.
func Parse(s string) (UUID, error) {
	var err error
	var id UUID
	var parsed uuid.UUID

	if parsed, err = uuid.Parse(s); err != nil {
		return id, err
	}

	return UUID{parsed}, nil
}

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
	var ok bool
	var str string

	if str, ok = v.(string); !ok {
		return fmt.Errorf("Value for unmarshaling was not a string: %v", v)
	}

	return u.UnmarshalJSON([]byte(str))
}

// MarshalGQL implements the graphql.Marshal interface.
func (u UUID) MarshalGQL(w io.Writer) {
	marshaled, _ := u.MarshalJSON()
	w.Write(marshaled)
}

// Size is required to implement the proto.Marshaler interface.
func (u *UUID) Size() int {
	if u == nil {
		return 0
	}
	return 16
}

// MarshalTo is required to implement the proto.Marshaler interface.
func (u *UUID) MarshalTo(data []byte) (int, error) {
	if u == nil {
		return 0, nil
	}
	copy(data, u.UUID[:])
	return 16, nil
}

// Unmarshal is required to implement the proto.Marshaler interface.
func (u *UUID) Unmarshal(data []byte) error {
	var err error
	var uid uuid.UUID

	if len(data) == 0 {
		u = nil
		return nil
	}

	if uid, err = uuid.FromBytes(data); err == nil {
		u.UUID = uid
	} else {
		return err
	}
}

// MarshalBSONValue implements the bson.ValueMarshaler interface.
func (u UUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	val := bsonx.Binary(0x04, u.UUID[:])
	return val.MarshalBSONValue()
}

// UnmarshalBSONValue implements the bson.ValueUnmarshaler interface.
func (u *UUID) UnmarshalBSONValue(bsonType bsontype.Type, data []byte) error {
	if bsonType != bsontype.Binary || data[0] != 0x10 || data[4] != 0x04 {
		return fmt.Errorf("Could not unmarshal %v as a UUID", bsonType)
	}

	return u.Unmarshal(data[5:])
}

// MarshalJSON implements the json.Marshaler interface.
func (u UUID) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(u.String())), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *UUID) UnmarshalJSON(data []byte) error {
	if parsed, err := uuid.Parse(string(data)); err == nil {
		u.UUID = parsed
	} else {
		return err
	}
}
