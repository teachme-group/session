package config

import (
	"github.com/Markuysa/pkg/consul"
	"github.com/Markuysa/pkg/log"
	"github.com/Markuysa/pkg/prober"
	promLoager "github.com/Markuysa/pkg/prometheus"
	"github.com/teachme-group/session/internal/service/client"

	"github.com/Markuysa/pkg/redis"
	"github.com/Markuysa/pkg/srv/grpc"
)

type Config struct {
	ClientSession client.Config     `yaml:"client_session"`
	Redis         redis.Config      `yaml:"redis"`
	Logger        log.Config        `yaml:"logger"`
	GRPC          grpc.Config       `yaml:"grpc"`
	Consul        consul.Config     `validate:"required" yaml:"consul"`
	Probes        prober.Config     `validate:"required" yaml:"probes"`
	Prometheus    promLoager.Config `validate:"required" yaml:"prometheus"`
}
