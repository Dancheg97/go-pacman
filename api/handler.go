package api

import (
	"context"

	"gitea.dancheg97.ru/dancheg97/go-pacman/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handlers struct {
	Packager *pkg.Packager
}

var ErrUnknown = status.Error(codes.NotFound, `unknown error`)

func (s Handlers) Add(ctx context.Context, in *pb.Packages) (*pb.Empty, error) {
	err := s.Packager.Add(in.Packages)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
