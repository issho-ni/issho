package graphql

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	log "github.com/sirupsen/logrus"
)

// MarshalTime implements graphql.Marshaler for time.Time.
func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		ts := t.Format(time.RFC3339)
		if _, err := w.Write([]byte(strconv.Quote(ts))); err != nil {
			log.Errorf("Error marshalling %v to GraphQL: %s", ts, err)
		}
	})
}

// UnmarshalTime implements graphql.Unmarshaler for time.Time.
func UnmarshalTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(time.RFC3339, v)
	case int:
		return time.Unix(int64(v), 0), nil
	default:
		return time.Now(), fmt.Errorf("Could not parse %v as time", v)
	}
}
