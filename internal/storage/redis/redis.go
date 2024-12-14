package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/teachme-group/session/pkg/errlist"

	v1 "github.com/teachme-group/session/pkg/api/grpc/v1"

	"github.com/Markuysa/pkg/tracer"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

const (
	sessionKeyfmt = "session:%s:%s"
)

type cache struct {
	rd *redis.Client
}

func NewStorage(rd *redis.Client) *cache {
	return &cache{rd: rd}
}

func (c *cache) SetClientSession(ctx context.Context, session *v1.Session) error {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	data, err := proto.Marshal(session)
	if err != nil {
		return err
	}

	key := fmt.Sprintf(sessionKeyfmt, session.GetAccessToken(), session.GetUserId())

	return c.rd.Set(
		ctx,
		key,
		data,
		time.Duration(session.ExpireIn),
	).Err()
}

func (c *cache) FindSessionByAccessToken(ctx context.Context, accessToken string) (*v1.Session, error) {
	ctx, span, _ := tracer.NewSpan(ctx)
	defer span.Finish()

	pattern := fmt.Sprintf("session:%s:*", accessToken)

	var cursor uint64
	var session *v1.Session

	for {
		keys, nextCursor, err := c.rd.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return nil, err
		}

		for _, key := range keys {
			result, err := c.rd.Get(ctx, key).Result()
			if err != nil {
				return nil, err
			}

			tempSession := &v1.Session{}
			if err := proto.Unmarshal([]byte(result), tempSession); err != nil {
				return nil, err
			}

			if tempSession.GetAccessToken() == accessToken {
				session = tempSession
				break
			}
		}

		if session != nil || nextCursor == 0 {
			break
		}

		cursor = nextCursor
	}

	if session == nil {
		return nil, errlist.ErrUnauthorized
	}

	return session, nil
}
