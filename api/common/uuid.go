package common

import (
	"fmt"
	"io"

	"github.com/issho-ni/issho/internal/pkg/uuid"

	log "github.com/sirupsen/logrus"
)

// ParseUUID decodes s into a UUID or returns an error.
func ParseUUID(s string) (id *UUID, err error) {
	if parsed, err := uuid.Parse(s); err == nil {
		id = &UUID{&parsed}
	}

	return
}

// NewUUID generates a new UUID.
func NewUUID() *UUID {
	id := uuid.New()
	return &UUID{&id}
}

// UnmarshalGQL implements the graphql.Unmarshal interface.
func (u *UUID) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("Value for unmarshalling was not a string: %v", v)
	}

	return u.Uuid.UnmarshalJSON([]byte(str))
}

// MarshalGQL implements the graphql.Marshal interface.
func (u UUID) MarshalGQL(w io.Writer) {
	marshaled, _ := u.Uuid.MarshalJSON()
	if _, err := w.Write(marshaled); err != nil {
		log.Errorf("Error marshalling %v to GraphQL: %s", u, err)
	}
}
