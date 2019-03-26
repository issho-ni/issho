package graphql

import (
	"io"
	"testing"
	"time"
)

type testWriter struct {
	io.Writer
	bytes []byte
}

func (w *testWriter) Write(p []byte) (int, error) {
	w.bytes = append(w.bytes, p...)
	return len(p), nil
}

func (w *testWriter) String() string {
	return string(w.bytes)
}

func TestMarshalTime(t *testing.T) {
	ts, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05-07:00")

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Time", args{ts}, "\"2006-01-02T15:04:05-07:00\""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &testWriter{}
			MarshalTime(tt.args.t).MarshalGQL(w)

			if got := w.String(); got != tt.want {
				t.Errorf("MarshalTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalTime(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"String", args{"2006-01-02T15:04:05-07:00"}, 1136239445, false},
		{"Integer", args{1136239445}, 1136239445, false},
		{"Bad type", args{false}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalTime(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Unix() != tt.want {
				t.Errorf("UnmarshalTime() = %v, want %v", got.Unix(), tt.want)
			}
		})
	}
}
