package grpc

import (
	context "context"

	"google.golang.org/grpc"
)

const (
	addr setting = "addr"
)

func init() {
	v.SetDefault(string(addr), "scoring")
}

type Conn struct {
	*grpc.ClientConn
}

func Connect(ctx context.Context) (*grpc.ClientConn, error) {
	connAddr := v.GetString(string(addr)) + ":" + v.GetString(string(port))
	return grpc.Dial(connAddr, grpc.WithInsecure())
}
