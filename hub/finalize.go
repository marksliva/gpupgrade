package hub

import (
	"context"
	"fmt"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/hashicorp/go-multierror"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/step"
)

func (s *Server) Finalize(_ *idl.FinalizeRequest, stream idl.CliToHub_FinalizeServer) (err error) {
	st, err := step.Begin(s.StateDir, "finalize", stream)
	if err != nil {
		return err
	}

	defer func() {
		if ferr := st.Finish(); ferr != nil {
			err = multierror.Append(err, ferr).ErrorOrNil()
		}

		if err != nil {
			gplog.Error(fmt.Sprintf("finalize: %s", err))
		}
	}()

	// This runner runs all commands against the target cluster.
	targetRunner := &greenplumRunner{
		masterPort:          s.Target.MasterPort(),
		masterDataDirectory: s.Target.MasterDataDir(),
		binDir:              s.Target.BinDir,
	}

	if s.Source.HasStandby() {
		st.Run(idl.Substep_FINALIZE_UPGRADE_STANDBY, func(streams step.OutStreams) error {
			// XXX this probably indicates a bad abstraction
			targetRunner.streams = streams

			// TODO: Persist the standby to config.json and update the
			//  source & target clusters.
			// todo: replace StandbyConfig with SegInfo and pass the TargetInitializeConfig.Standby directly in
			standby := s.TargetInitializeConfig.Standby
			return UpgradeStandby(targetRunner, StandbyConfig{
				Port:          standby.Port,
				Hostname:      standby.Hostname,
				DataDirectory: standby.DataDir,
			})
		})
	}

	if s.Source.HasMirrors() {
		st.Run(idl.Substep_FINALIZE_UPGRADE_MIRRORS, func(streams step.OutStreams) error {
			// XXX this probably indicates a bad abstraction
			targetRunner.streams = streams

			return UpgradeMirrors(s.StateDir, s.Target.MasterPort(), &s.TargetInitializeConfig, targetRunner)
		})
	}

	st.Run(idl.Substep_FINALIZE_SHUTDOWN_TARGET_CLUSTER, func(streams step.OutStreams) error {
		return StopCluster(streams, s.Target, false)
	})

	st.Run(idl.Substep_FINALIZE_UPDATE_TARGET_CATALOG_AND_CLUSTER_CONFIG, func(streams step.OutStreams) error {
		return s.UpdateCatalogAndClusterConfig(streams)
	})

	st.Run(idl.Substep_FINALIZE_RENAME_DATA_DIRECTORIES, func(_ step.OutStreams) error {
		return s.RenameDataDirectories()
	})

	st.Run(idl.Substep_FINALIZE_UPDATE_TARGET_CONF_FILES, func(_ step.OutStreams) error {
		return s.UpdateConfFiles()
	})

	if s.Source.HasMirrors() {
		st.Run(idl.Substep_FINALIZE_UPDATE_RECOVERY_CONFS, func(streams step.OutStreams) error {
			return UpdateRecoveryConfs(context.Background(), s.agentConns, s.Source, s.Target, s.TargetInitializeConfig)
		})
	}

	st.Run(idl.Substep_FINALIZE_START_TARGET_CLUSTER, func(streams step.OutStreams) error {
		return StartCluster(streams, s.Target, false)
	})

	return st.Err()
}
