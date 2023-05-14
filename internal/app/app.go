package app

import (
	"github.com/iamasmith/skeleton/internal/config"
	"github.com/iamasmith/skeleton/internal/logging"
	"github.com/iamasmith/skeleton/internal/server"
	"github.com/iamasmith/skeleton/internal/version"
	"go.uber.org/zap"
)

type AppState struct {
	logger *zap.SugaredLogger
	s      *server.ServerState
}

func Setup() (*server.ServerState, *AppState) {
	app := AppState{logger: logging.Setup(config.Config.LogLevel)}
	app.s = server.New(config.Config.ListenBind)
	app.logger.Debug("ServerState created")

	// Add handlers with mux here
	// mux := s.Mux()
	app.logger.Infof("Starting %s %s %s (Build: %s, Built: %s)", version.Name(), version.Version(), version.BuildType(), version.BuildId(), version.BuildDate())
	app.logger.Infof("Server bound to %s", config.Config.ListenBind)
	return app.s, &app
}

func (s *AppState) Stop() {
	s.logger.Debug("App stopped")
}
