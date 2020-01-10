package common

import (
	"bytes"
	fmt "fmt"
	"testing"

	"github.com/issho-ni/issho/internal/pkg/uuid"

	guuid "github.com/google/uuid"
)

var uuidString = "4915e120-3594-b29e-bd92-62abff23e1c6"
var uuidBytes = [16]byte{0x49, 0x15, 0xe1, 0x20, 0x35, 0x94, 0xb2, 0x9e, 0xbd, 0x92, 0x62, 0xab, 0xff, 0x23, 0xe1, 0xc6}

func TestUUID_UnmarshalGQL(t *testing.T) {
	type fields struct {
		UUID *uuid.UUID
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Valid UUID", fields{&uuid.UUID{UUID: guuid.UUID(uuidBytes)}}, args{uuidString}, false},
		{"Invalid UUID", fields{}, args{""}, true},
		{"Non-string", fields{}, args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UUID{
				Uuid: tt.fields.UUID,
			}
			if err := u.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("UUID.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUUID_MarshalGQL(t *testing.T) {
	type fields struct {
		UUID *uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		wantW  string
	}{
		{"Valid UUID", fields{&uuid.UUID{UUID: guuid.UUID(uuidBytes)}}, fmt.Sprintf("\"%s\"", uuidString)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UUID{
				Uuid: tt.fields.UUID,
			}
			w := &bytes.Buffer{}
			u.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("UUID.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
