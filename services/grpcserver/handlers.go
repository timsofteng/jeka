package grpcserver

import (
	"context"
	"fmt"

	"github.com/timsofteng/jeka/services/grpcserver/pb"
)

func (s *server) GetRandomText(
	ctx context.Context,
	_ *pb.GetRandomTextRequest,
) (*pb.GetRandomTextResponse, error) {
	resp, err := s.Srv.Text.Rand(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to call rand text service: %w", err)
	}

	return &pb.GetRandomTextResponse{
		Text: resp.Text,
	}, nil
}

func (s *server) GetRandomImg(
	ctx context.Context,
	_ *pb.GetRandomImgRequest,
) (*pb.GetRandomImgResponse, error) {
	resp, err := s.Srv.Image.RandImg(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to call rand text service: %w", err)
	}

	return &pb.GetRandomImgResponse{
		Url: resp.URL,
	}, nil
}
