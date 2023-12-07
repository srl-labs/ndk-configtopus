package configtopus

import (
	"context"
	"encoding/json"

	"github.com/nokia/srlinux-ndk-go/ndk"
)

const greeterKeyPath = ".greeter"

// updateState updates the state of the application.

func (a *App) updateState(ctx context.Context) {
	jsData, err := json.Marshal(a.configState)
	if err != nil {
		a.logger.Info().Msgf("failed to marshal json data: %v", err)
		return
	}

	a.telemetryAddOrUpdate(ctx, greeterKeyPath, string(jsData))
}

// telemetryAddOrUpdate updates the state of the application using provided path and data.

func (a *App) telemetryAddOrUpdate(ctx context.Context, jsPath string, jsData string) {
	a.logger.Info().Msgf("updating: %s: %s", jsPath, jsData)

	key := &ndk.TelemetryKey{JsPath: jsPath}
	data := &ndk.TelemetryData{JsonContent: jsData}
	info := &ndk.TelemetryInfo{Key: key, Data: data}
	req := &ndk.TelemetryUpdateRequest{
		State: []*ndk.TelemetryInfo{info},
	}

	a.logger.Info().Msgf("Telemetry Request: %+v", req)

	r1, err := a.TelemetryServiceClient.TelemetryAddOrUpdate(ctx, req)
	if err != nil {
		a.logger.Info().Msgf("Could not update telemetry key=%s: err=%v", jsPath, err)
		return
	}

	a.logger.Info().Msgf("Telemetry add/update status: %s, error_string: %q",
		r1.GetStatus().String(), r1.GetErrorStr())
}
