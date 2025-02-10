package link

import (
	"context"

	linkv2 "github.com/Maksim646/protos/gen/go/link"
	"google.golang.org/grpc"
)

type serverAPI struct {
	linkv2.UnimplementedLinkServer
}

func Register(gRPC *grpc.Server) {
	linkv2.RegisterLinkServer(gRPC, &serverAPI{})
}

func (s *serverAPI) CreateShorterLink(ctx context.Context, req *linkv2.CreateShorterLinkRequest) (*linkv2.CreateShorterLinkResponse, error) {

	panic("implement me")
}

func (s *serverAPI) GetOriginalLink(ctx context.Context, req *linkv2.GetOriginalLinkRequest) (*linkv2.GetOriginalLinkResponse, error) {

	panic("implement me")
}
