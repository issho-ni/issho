package service

import (
	"os"
	"reflect"
	"testing"
)

func TestNewServerConfig(t *testing.T) {
	type args struct {
		name        string
		defaultPort string
	}
	tests := []struct {
		name string
		args args
		envs map[string]string
		want *ServerConfig
	}{
		{"Defaults", args{"default", "1234"}, map[string]string{}, &ServerConfig{"default", "1234", "localhost+2.pem", "localhost+2-key.pem"}},
		{"Port from env", args{"port", "1234"}, map[string]string{"PORT_PORT": "3456"}, &ServerConfig{"port", "3456", "localhost+2.pem", "localhost+2-key.pem"}},
		{"TLS cert/key from env", args{"tls", "1234"}, map[string]string{"TLS_CERT": "cert.pem", "TLS_KEY": "key.pem"}, &ServerConfig{"tls", "1234", "cert.pem", "key.pem"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envs {
				key := k
				os.Setenv(k, v)
				defer func() { os.Unsetenv(key) }()
			}

			if got := NewServerConfig(tt.args.name, tt.args.defaultPort); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
