package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/teachme-group/session/internal/config"
	clientRepos "github.com/teachme-group/session/internal/repository/client"
	clientService "github.com/teachme-group/session/internal/service/client"
	"github.com/teachme-group/session/internal/transport/client"
	closerPkg "github.com/teachme-group/session/pkg/closer"

	"github.com/Markuysa/pkg/logger"
	"github.com/Markuysa/pkg/redis"
	"github.com/Markuysa/pkg/srv/grpc"
)

func Run(_ context.Context, cfg *config.Config) error {
	closer := closerPkg.New()

	err := logger.InitLogger(cfg.Logger)
	if err != nil {
		return err
	}

	rdConn, err := redis.New(cfg.Redis)
	if err != nil {
		return err
	}
	closer.AddErrCloser(rdConn.Close)

	sessionsRepository := clientRepos.New(rdConn)
	sessionService := clientService.New(
		cfg.ClientSession,
		sessionsRepository,
	)

	grpcTransport := client.New(sessionService)

	grpc, err := grpc.NewServer(
		grpc.WithConfig(&cfg.GRPC),
		grpc.WithRegistes(
			grpcTransport,
		),
	)
	if err != nil {
		return err
	}
	closer.AddCloser(grpc.GracefulStop)

	logger.Logger.Info("started app")

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quitCh

	return closer.Close()
}
