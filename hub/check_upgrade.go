package hub

import (
	"sync"

	"github.com/pkg/errors"

	"github.com/hashicorp/go-multierror"

	"github.com/greenplum-db/gpupgrade/step"
)

func (h *Hub) CheckUpgrade(stream step.OutStreams) error {
	var wg sync.WaitGroup
	checkErrs := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()

		stateDir := h.StateDir
		err := UpgradeMaster(h.Source, h.Target, stateDir, stream, true, false)
		if err != nil {
			checkErrs <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		agentConns, agentConnsErr := h.AgentConns()

		if agentConnsErr != nil {
			checkErrs <- errors.Wrap(agentConnsErr, "failed to connect to gpupgrade agent")
		}

		dataDirPairMap, dataDirPairsErr := h.GetDataDirPairs()

		if dataDirPairsErr != nil {
			checkErrs <- errors.Wrap(dataDirPairsErr, "failed to get old and new primary data directories")
		}

		upgradeErr := UpgradePrimaries(true, "", agentConns, dataDirPairMap, h.Source, h.Target, h.UseLinkMode)

		if upgradeErr != nil {
			checkErrs <- upgradeErr
		}
	}()

	wg.Wait()
	close(checkErrs)

	var multiErr *multierror.Error
	for err := range checkErrs {
		multiErr = multierror.Append(multiErr, err)
	}

	return multiErr.ErrorOrNil()
}
