package client

type service struct {
	cfg   Config
	repos repository
}

func New(
	cfg Config,
	repos repository,
) *service {
	return &service{
		cfg:   cfg,
		repos: repos,
	}
}
