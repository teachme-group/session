package client

import (
	"context"

	v1 "github.com/teachme-group/session/pkg/api/grpc/v1"
)

type (
	service interface {
		// CreateSession creates a new session.
		CreateSession(ctx context.Context, req *v1.ClientSetSessionRequest) (*v1.ClientSetSessionResponse, error)
		// ClientAuthorize authorizes a client.
		ClientAuthorize(ctx context.Context, req *v1.ClientAuthRequest) (*v1.ClientAuthResponse, error)
	}
)
