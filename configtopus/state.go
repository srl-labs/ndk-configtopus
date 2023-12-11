package configtopus

import "github.com/openconfig/ygot/ygot"

func (a *App) syncState() error {
	err := a.NDKAgent.DeleteState()
	if err != nil {
		a.logger.Error().Err(err).Msg("Failed to delete state")
	}

	b, err := ygot.Marshal7951(a.configState)
	if err != nil {
		a.logger.Error().Err(err).Msg("Failed to marshal config to JSON")
	}
	a.logger.Debug().Msgf("Loaded config in JSON: %s", b)

	sb := string(b)

	err = a.updateState("", sb)
	if err != nil {
		a.logger.Error().Err(err).Msg("Failed to update state")
		return err
	}

	return nil
}

func (a *App) updateState(key, data string) error {
	err := a.NDKAgent.UpdateState(key, data)
	if err != nil {
		a.logger.Error().Err(err).Msg("Failed to update state")

		return err
	}

	return nil
}
