package main

import (
	"context"
	"net/http"

	pb "github.com/tsub/grpc-go-sandbox/protobuf"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	endpoint = "localhost:50051"
)

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
    mux := runtime.NewServeMux(opts...)
    dialOpts := []grpc.DialOption{grpc.WithInsecure()}
    conn, err := grpc.Dial(endpoint, dialOpts...)
    if err != nil {
        return nil, err
    }
    err = pb.RegisterGreeterHandler(ctx, mux, conn)
    if err != nil {
        return nil, err
    }

    return mux, nil
}

func run(address string, opts ...runtime.ServeMuxOption) error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    gw, err := newGateway(ctx, opts...)
    if err != nil {
        return err
    }

    return http.ListenAndServe(address, gw)
}

func main() {
    if err := run(":3000"); err != nil {
        panic(err)
    }
}
