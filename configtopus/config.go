package configtopus

import "github.com/srl-labs/ndk-configtopus/configtopus/config"

func (a *App) loadConfig() {
	err := config.Unmarshal(a.NDKAgent.Config, a.configState)
	if err != nil {
		a.logger.Error().Err(err).Msg("Failed to unmarshal config")
	}

	a.logger.Info().Msgf("Loaded config: %+v", a.configState)
}
