package hub

import (
	"fmt"
	"path/filepath"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/hashicorp/go-multierror"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/step"
)

func (h *Hub) Execute(request *idl.ExecuteRequest, stream idl.CliToHub_ExecuteServer) (err error) {
	masterBackupDir := filepath.Join(h.StateDir, "master.bak")

	s, err := BeginStep(h.StateDir, "execute", stream)
	if err != nil {
		return err
	}

	defer func() {
		if ferr := s.Finish(); ferr != nil {
			err = multierror.Append(err, ferr).ErrorOrNil()
		}

		if err != nil {
			gplog.Error(fmt.Sprintf("execute: %s", err))
		}
	}()

	s.Run(idl.Substep_UPGRADE_MASTER, func(streams step.OutStreams) error {
		stateDir := h.StateDir
		return UpgradeMaster(h.Source, h.Target, stateDir, streams, false, h.UseLinkMode)
	})

	s.Run(idl.Substep_COPY_MASTER, func(streams step.OutStreams) error {
		return h.CopyMasterDataDir(streams, masterBackupDir)
	})

	s.Run(idl.Substep_UPGRADE_PRIMARIES, func(_ step.OutStreams) error {
		return h.UpgradePrimaries(false, masterBackupDir)
	})

	s.Run(idl.Substep_START_TARGET_CLUSTER, func(streams step.OutStreams) error {
		return StartCluster(streams, h.Target)
	})

	return s.Err()
}
