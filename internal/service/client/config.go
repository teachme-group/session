package client

import "time"

type Config struct {
	SessionExpireTime time.Duration `yaml:"session_expire_time" default:"30m"`
}
