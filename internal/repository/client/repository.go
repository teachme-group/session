package client

import (
	"context"

	rd "github.com/teachme-group/session/internal/storage/redis"
	v1 "github.com/teachme-group/session/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/tracer"
	"github.com/redis/go-redis/v9"
)

type repo struct {
	storage storage
}

func New(client *redis.Client) *repo {
	return &repo{
		storage: rd.NewStorage(client),
	}
}

func (r *repo) SetClientSession(ctx context.Context, session *v1.Session) error {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return r.storage.SetClientSession(ctx, session)
}

func (r *repo) FindSessionByAccessToken(ctx context.Context, accessToken string) (*v1.Session, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	return r.storage.FindSessionByAccessToken(ctx, accessToken)
}
