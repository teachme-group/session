package config

import (
	"github.com/teachme-group/session/internal/service/client"

	"github.com/Markuysa/pkg/logger"
	"github.com/Markuysa/pkg/redis"
	"github.com/Markuysa/pkg/srv/grpc"
)

type Config struct {
	ClientSession client.Config `yaml:"client_session"`
	Redis         redis.Config  `yaml:"redis"`
	Logger        logger.Config `yaml:"logger"`
	GRPC          grpc.Config   `yaml:"grpc"`
}
