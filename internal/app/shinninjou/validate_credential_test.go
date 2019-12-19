package shinninjou

import (
	"context"
	"reflect"
	"testing"

	"github.com/issho-ni/issho/api/shinninjou"
	"github.com/issho-ni/issho/internal/pkg/grpc"
	"github.com/issho-ni/issho/internal/pkg/mongo"
)

func TestServer_ValidateCredential(t *testing.T) {
	type fields struct {
		Client mongo.Client
	}
	type args struct {
		ctx context.Context
		in  *shinninjou.Credential
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse *shinninjou.CredentialResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Server: &grpc.Server{
					MongoClient: tt.fields.Client,
				},
			}
			gotResponse, err := s.ValidateCredential(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ValidateCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Server.ValidateCredential() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
