package hub

import (
	"fmt"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/hashicorp/go-multierror"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/step"
)

func (s *Server) Revert(_ *idl.RevertRequest, stream idl.CliToHub_RevertServer) (err error) {
	st, err := step.Begin(s.StateDir, "revert", stream)
	if err != nil {
		return err
	}

	defer func() {
		if ferr := st.Finish(); ferr != nil {
			err = multierror.Append(err, ferr).ErrorOrNil()
		}

		if err != nil {
			gplog.Error(fmt.Sprintf("revert: %s", err))
		}
	}()

	err = DeleteSegmentAndStandbyDirectories(s.agentConns, s.Config.Target)
	if err != nil {
		return err
	}

	err = s.StopAgents()
	if err != nil {
		return err
	}

	err = DeleteMasterDataDirectory(s.Config.Target.MasterDataDir())
	if err != nil {
		return err
	}

	return st.Err()
}
