package client

import (
	"context"

	v1 "github.com/teachme-group/session/pkg/api/grpc/v1"
)

type repository interface {
	SetClientSession(ctx context.Context, session *v1.Session) error
	FindSessionByAccessToken(ctx context.Context, accessToken string) (*v1.Session, error)
}
