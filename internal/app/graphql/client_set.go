package graphql

import (
	"reflect"

	"github.com/issho-ni/issho/api/kazoku"
	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/shinninjou"
	"github.com/issho-ni/issho/api/youji"
	"github.com/issho-ni/issho/internal/pkg/grpc"
)

// ClientSet is a basket for all GRPC service clients.
type ClientSet interface {
	AllClients() []grpc.Client
}

type clientSet struct {
	KazokuClient     *kazoku.Client
	NinkaClient      *ninka.Client
	NinshouClient    *ninshou.Client
	ShinninjouClient *shinninjou.Client
	YoujiClient      *youji.Client
}

// AllClients returns a slice of all configured clients.
func (cs *clientSet) AllClients() (clients []grpc.Client) {
	v := reflect.Indirect(reflect.ValueOf(cs))

	for i := 0; i < v.NumField(); i++ {
		if client := reflect.Indirect(reflect.ValueOf(v.Field(i).Interface())); client.IsValid() {
			grpcClient := client.FieldByName("Client").Interface().(grpc.Client)
			clients = append(clients, grpcClient)
		}
	}

	return
}
