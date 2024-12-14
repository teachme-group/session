package client

import (
	"context"

	v1 "github.com/teachme-group/session/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/tracer"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *service) CreateSession(ctx context.Context, req *v1.ClientSetSessionRequest) (*v1.ClientSetSessionResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	session := &v1.Session{
		Id:           uuid.NewString(),
		UserId:       req.ClientId,
		ExpireIn:     s.cfg.SessionExpireTime.Milliseconds(),
		AccessToken:  uuid.NewString(),
		RefreshToken: uuid.NewString(),
		SignedAt:     timestamppb.Now(),
	}

	err := s.repos.SetClientSession(ctx, session)
	if err != nil {
		return nil, err
	}

	return &v1.ClientSetSessionResponse{
		AccessToken: session.GetAccessToken(),
	}, nil
}

func (s *service) ClientAuthorize(ctx context.Context, req *v1.ClientAuthRequest) (*v1.ClientAuthResponse, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	session, err := s.repos.FindSessionByAccessToken(ctx, req.GetClientSecret())
	if err != nil {
		return nil, err
	}

	return &v1.ClientAuthResponse{
		Id:           session.GetUserId(),
		UserId:       session.GetUserId(),
		AccessToken:  session.GetAccessToken(),
		RefreshToken: session.GetRefreshToken(),
		SignedAt:     session.GetSignedAt(),
		UpdatedAt:    session.GetUpdatedAt(),
		ExpireIn:     session.GetExpireIn(),
	}, nil
}
