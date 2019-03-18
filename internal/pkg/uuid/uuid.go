package uuid

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx"
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
	marshaled, _ := u.MarshalJSON()
	w.Write(marshaled)
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

// MarshalBSONValue implements the bson.ValueMarshaler interface.
func (u UUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	val := bsonx.Binary(0x04, u.UUID[:])
	return val.MarshalBSONValue()
}

// UnmarshalBSONValue implements the bson.ValueUnmarshaler interface.
func (u *UUID) UnmarshalBSONValue(bsonType bsontype.Type, data []byte) error {
	if bsonType != bsontype.Binary && data[4] != 0x04 {
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
	return u.Unmarshal(data)
}
