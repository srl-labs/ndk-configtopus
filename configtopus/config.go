package configtopus

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/srl-labs/ndk-configtopus/configtopus/config"
)

func (a *App) loadConfig() {
	// clear the configState since it might contain old data
	// and we always load the full config
	// retrieved via gNMI.Get.
	a.configState = &config.App_Configtopus{}

	if a.NDKAgent.Config != nil {
		err := config.Unmarshal(a.NDKAgent.Config, a.configState)
		if err != nil {
			a.logger.Error().Err(err).Msg("Failed to unmarshal config")
		}
	}

	a.logger.Debug().Msgf("Loaded config: %s", spew.Sdump(a.configState))
}
