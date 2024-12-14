package client

import (
	"context"
	"errors"

	v1 "github.com/teachme-group/session/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/errs"
	"github.com/Markuysa/pkg/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type sessionTransport struct {
	service service
	v1.UnimplementedSessionServiceServer
}

func New(s service) *sessionTransport {
	return &sessionTransport{service: s}
}

func (s *sessionTransport) RegisterServer(server *grpc.Server) {
	v1.RegisterSessionServiceServer(server, s)
}

func (s *sessionTransport) ClientAuth(ctx context.Context, req *v1.ClientAuthRequest) (*v1.ClientAuthResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	resp, err := s.service.ClientAuthorize(ctx, req)
	if err != nil {
		// TODO вынести в пакет ошибок.
		customErr := &errs.Error{}
		if errors.As(err, &customErr) {
			return nil, status.Error(codes.Code(customErr.Code), customErr.Msg)
		}

		return resp, err
	}

	return resp, nil
}

func (s *sessionTransport) ClientSetSession(ctx context.Context, req *v1.ClientSetSessionRequest) (*v1.ClientSetSessionResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return s.service.CreateSession(ctx, req)
}
