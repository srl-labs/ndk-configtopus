package configtopus

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/srl-labs/bond"
)

const (
	ndkSocket            = "unix:///opt/srlinux/var/run/sr_sdk_service_manager:50053"
	grpcServerUnixSocket = "unix:///opt/srlinux/var/run/sr_gnmi_server"
)

// App is the greeter application struct.

type App struct {
	Name string

	// configState holds the application configuration and state.
	// configState *ConfigState

	NDKAgent *bond.Agent

	logger *zerolog.Logger
}

// New creates a new App instance and connects to NDK socket.
// It also creates the NDK service clients and registers the agent with NDK.

func New(name string, logger *zerolog.Logger, agent *bond.Agent) *App {
	return &App{
		Name: name,

		// configState: &ConfigState{},

		logger: logger,

		NDKAgent: agent,
	}
}

// Start starts the application.

func (a *App) Start(ctx context.Context) {
	// go a.receiveConfigNotifications(ctx)

	for {
		select {
		case <-a.NDKAgent.ConfigReceivedCh:
			a.logger.Info().Msg("Received full config")

			// a.processConfig(ctx)

			// a.updateState(ctx)

		case <-ctx.Done():
			a.stop()
			return
		}
	}
}

// stop exits the application gracefully.

func (a *App) stop() {
	a.logger.Info().Msg("Got a signal to exit, unregistering configtopus agent, bye!")

	// unregister agent
	_ = a.NDKAgent.Stop()

	a.logger.Info().Msg("Configtopus unregistered successfully!")
}
