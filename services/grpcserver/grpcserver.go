package grpcserver

import (
	"fmt"
	"net"

	"github.com/timsofteng/jeka/services/grpcserver/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/timsofteng/jeka/services/images"
	"github.com/timsofteng/jeka/services/text"
	"github.com/timsofteng/jeka/services/video"
	"github.com/timsofteng/jeka/services/voice"
)

type GRPCServer struct {
	server *grpc.Server
}

type server struct {
	// type embedded to comply with Google lib
	pb.UnimplementedJekaServer
	Srv Services
}

type Services struct {
	Video *video.Video
	Image *images.Images
	Text  *text.Text
	Voice *voice.Voice
}

func New(srv Services) *GRPCServer {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterJekaServer(grpcServer, &server{
		UnimplementedJekaServer: pb.UnimplementedJekaServer{},
		Srv:                     srv,
	})

	return &GRPCServer{server: grpcServer}
}

func (g *GRPCServer) Start(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen port %s: %w", port, err)
	}

	if err := g.server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve grpc server: %w", err)
	}

	return nil
}

func (g *GRPCServer) Stop() {
	g.server.GracefulStop()
}
