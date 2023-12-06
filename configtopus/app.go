// --8<-- [start:pkg-greeter]
package configtopus

// --8<-- [end:pkg-greeter]

import (
	"context"
	"time"

	"github.com/nokia/srlinux-ndk-go/ndk"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// --8<-- [start:pkg-greeter-const].
const (
	ndkSocket            = "unix:///opt/srlinux/var/run/sr_sdk_service_manager:50053"
	grpcServerUnixSocket = "unix:///opt/srlinux/var/run/sr_gnmi_server"
	AppName              = "configtopus"
)

// --8<-- [end:pkg-greeter-const]

// App is the greeter application struct.
// --8<-- [start:app-struct].
type App struct {
	Name  string
	AppID uint32

	// configState holds the application configuration and state.
	configState *ConfigState
	// configReceivedCh chan receives the value when the full config
	// is received by the stream client.
	configReceivedCh chan struct{}

	gRPCConn     *grpc.ClientConn
	logger       *zerolog.Logger
	retryTimeout time.Duration

	// NDK Service clients
	SDKMgrServiceClient       ndk.SdkMgrServiceClient
	NotificationServiceClient ndk.SdkNotificationServiceClient
	TelemetryServiceClient    ndk.SdkMgrTelemetryServiceClient
}

// --8<-- [end:app-struct]

// New creates a new App instance and connects to NDK socket.
// It also creates the NDK service clients and registers the agent with NDK.
// --8<-- [start:new-app].
func New(ctx context.Context, logger *zerolog.Logger) *App {
	// connect to NDK socket
	conn, err := connect(ctx, ndkSocket)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("gRPC connect failed")
	}

	// --8<-- [start:create-ndk-clients]
	sdkMgrClient := ndk.NewSdkMgrServiceClient(conn)
	notifSvcClient := ndk.NewSdkNotificationServiceClient(conn)
	telemetrySvcClient := ndk.NewSdkMgrTelemetryServiceClient(conn)
	// --8<-- [end:create-ndk-clients]

	// register agent
	// http://learn.srlinux.dev/ndk/guide/dev/go/#register-the-agent-with-the-ndk-manager
	// --8<-- [start:register-agent]
	r, err := sdkMgrClient.AgentRegister(ctx, &ndk.AgentRegistrationRequest{})
	if err != nil || r.Status != ndk.SdkMgrStatus_kSdkMgrSuccess {
		logger.Fatal().
			Err(err).
			Str("status", r.GetStatus().String()).
			Msg("Agent registration failed")
	}
	// --8<-- [end:register-agent]

	logger.Info().
		Uint32("app-id", r.GetAppId()).
		Str("name", AppName).
		Msg("Application registered successfully!")

	// --8<-- [start:return-app]
	return &App{
		Name:  AppName,
		AppID: r.GetAppId(), //(1)!

		configState:      &ConfigState{},
		configReceivedCh: make(chan struct{}),

		logger:       logger,
		retryTimeout: 5 * time.Second,
		gRPCConn:     conn,

		SDKMgrServiceClient:       sdkMgrClient,
		NotificationServiceClient: notifSvcClient,
		TelemetryServiceClient:    telemetrySvcClient,
	}
	// --8<-- [end:return-app]
}

// --8<-- [end:new-app]

// Start starts the application.
// --8<-- [start:app-start].
func (a *App) Start(ctx context.Context) {
	go a.receiveConfigNotifications(ctx)

	for {
		select {
		case <-a.configReceivedCh:
			a.logger.Info().Msg("Received full config")

			// a.processConfig(ctx)

			// a.updateState(ctx)

		case <-ctx.Done():
			a.stop()
			return
		}
	}
}

// --8<-- [end:app-start]

// stop exits the application gracefully.
// --8<-- [start:app-stop].
func (a *App) stop() {
	a.logger.Info().Msg("Got a signal to exit, unregistering configtopus agent, bye!")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx = metadata.AppendToOutgoingContext(ctx, "agent_name", AppName)
	defer cancel()

	// unregister agent
	r, err := a.SDKMgrServiceClient.AgentUnRegister(ctx, &ndk.AgentRegistrationRequest{})
	if err != nil || r.Status != ndk.SdkMgrStatus_kSdkMgrSuccess {
		a.logger.Error().
			Err(err).
			Str("status", r.GetStatus().String()).
			Msgf("Agent unregistration failed %s", r.GetErrorStr())

		return
	}

	err = a.gRPCConn.Close()
	if err != nil {
		a.logger.Error().Err(err).Msg("Closing gRPC connection to NDK server failed")
	}

	a.logger.Info().Msg("Configtopus unregistered successfully!")
}

// --8<-- [end:app-stop]

// connect attempts connecting to the NDK socket with backoff and retry.
// https://learn.srlinux.dev/ndk/guide/dev/go/#establish-grpc-channel-with-ndk-manager-and-instantiate-an-ndk-client
// --8<-- [start:connect].
func connect(ctx context.Context, socket string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(ndkSocket,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	return conn, err
}

// --8<-- [end:connect]
