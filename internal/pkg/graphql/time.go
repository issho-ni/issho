package graphql

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalTime implements graphql.Marshaler for time.Time.
func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		ts := t.Format(time.RFC3339)
		w.Write([]byte(strconv.Quote(ts)))
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
