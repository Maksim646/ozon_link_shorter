package link

import (
	linkv2 "github.com/Maksim646/protos/gen/go/link"
	"google.golang.org/grpc"
)

type serverAPI struct {
	linkv2.UnimplementedLinkServer
}

func Link(gRPC *grpc.Server) {
	linkv2.RegisterLinkServer(gRPC, &serverAPI{})
}
