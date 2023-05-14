package app

import (
	"github.com/iamasmith/skeleton/internal/logging"
	"github.com/iamasmith/skeleton/internal/server"
	"github.com/iamasmith/skeleton/internal/version"
	"go.uber.org/zap"
)

type AppState struct {
	logger *zap.SugaredLogger
	s      *server.ServerState
}

func Setup() *server.ServerState {
	app := AppState{logger: logging.Setup()}
	app.s = &server.ServerState{}

	// Add handlers with mux here
	// mux := s.Mux()
	app.logger.Infof("Starting %s %s %s (Build: %s, Built: %s)", version.Name(), version.Version(), version.BuildType(), version.BuildId(), version.BuildDate())
	return app.s
}
